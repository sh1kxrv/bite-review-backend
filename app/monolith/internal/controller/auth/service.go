package auth

import (
	"bitereview/internal/config"
	"context"
	"shared/database/mongodb/entity"
	"shared/database/mongodb/repository"
	"shared/enum"
	"shared/errors"
	"shared/transfer/dto"
	"shared/transfer/ro"
	"shared/utils"
	"shared/utils/crypto"
	"shared/utils/helper"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (h *AuthService) createJwt(user *entity.User, key string, duration time.Duration) (string, time.Time, error) {
	signingKey := []byte(key)

	now := time.Now()
	expiresAt := now.Add(time.Duration(time.Hour * 24 * duration))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.JwtClaims{
		Role: user.Role,
		ID:   user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "bitereview",
		},
	})

	ss, err := token.SignedString(signingKey)
	return ss, expiresAt, err
}

func (h *AuthService) createJwtPair(user *entity.User) (ro.JwtPair, error) {
	accessToken, expAccess, accessErr := h.createJwt(user, config.C.Jwt.Secret, time.Duration(config.C.Jwt.Expire))
	if accessErr != nil {
		return ro.JwtPair{}, accessErr
	}
	refreshToken, expRefresh, refreshErr := h.createJwt(user, config.C.Jwt.RefreshSecret, time.Duration(config.C.Jwt.RefreshExpire))
	if refreshErr != nil {
		return ro.JwtPair{}, refreshErr
	}

	expRefreshFormatted := expRefresh.Format(time.RFC3339)
	expAccessFormatted := expAccess.Format(time.RFC3339)

	return ro.JwtPair{
		AccessToken:      accessToken,
		RefreshToken:     &refreshToken,
		AccessExpiresIn:  expAccessFormatted,
		RefreshExpiresIn: &expRefreshFormatted,
	}, nil
}

func (h *AuthService) Login(email, password string) (*ro.JwtPair, *helper.ServiceError) {
	withTimeout, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	user, err := h.userRepo.FindByEmail(withTimeout, email)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.RepositoryError)
	}

	if user == nil {
		return nil, helper.NewServiceError(err, errors.Forbidden)
	}

	if !crypto.CheckPasswordHash(password, user.Password) {
		return nil, helper.NewServiceError(err, errors.Forbidden)
	}

	pair, err := h.createJwtPair(user)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.JwtPairGenerationError)
	}

	go h.userRepo.UpdateLastSeen(context.TODO(), user.ID)

	return &pair, nil
}

func (h *AuthService) Register(authData *dto.AuthDataRegister) (*ro.JwtPair, *helper.ServiceError) {
	ctx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	_, err := h.userRepo.FindByEmail(ctx, authData.Email)
	if err == nil {
		return nil, helper.NewServiceError(err, errors.EntityAlreadyExists)
	}

	hashedPwd, err := crypto.HashPassword(authData.Password)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.CryptoError)
	}

	user, err := h.userRepo.CreateEntity(ctx, &entity.User{
		ID:       primitive.NewObjectID(),
		Email:    authData.Email,
		Password: hashedPwd,
		Role:     enum.RoleCritic,
		LastSeen: time.Now(),
	})

	if err != nil {
		return nil, helper.NewServiceError(err, errors.MakeRepositoryError("User"))
	}

	pair, err := h.createJwtPair(user)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.JwtPairGenerationError)
	}

	return &pair, nil
}

// FIXME: Добавить проверку JWT AccessToken'а в Header's с игнором exp. времени
func (h *AuthService) Refresh(refreshToken string) (*ro.JwtPair, *helper.ServiceError) {
	parsed, err := utils.ParseJwtToken(refreshToken, config.C.Jwt.RefreshSecret)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.JwtRefreshTokenInvalid)
	}

	parsedId, err := primitive.ObjectIDFromHex(parsed.ID)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.ParseIDError)
	}

	ctx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	user, err := h.userRepo.GetEntityByID(ctx, parsedId)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.EntityNotExists)
	}

	accessToken, expAccess, accessErr := h.createJwt(user, config.C.Jwt.Secret, time.Duration(config.C.Jwt.Expire))
	if accessErr != nil {
		return nil, helper.NewServiceError(err, errors.JwtPairGenerationError)
	}

	return &ro.JwtPair{
		AccessToken:      accessToken,
		AccessExpiresIn:  expAccess.Format(time.RFC3339),
		RefreshToken:     nil,
		RefreshExpiresIn: nil,
	}, nil
}
