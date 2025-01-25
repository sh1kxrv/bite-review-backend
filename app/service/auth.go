package service

import (
	"bitereview/app/config"
	"bitereview/app/crypto"
	"bitereview/app/enum"
	"bitereview/app/errors"
	"bitereview/app/helper"
	"bitereview/app/repository"
	"bitereview/app/schema"
	"bitereview/app/serializer"
	"bitereview/app/utils"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

type JwtPair struct {
	AccessToken      string  `json:"accessToken"`
	RefreshToken     *string `json:"refreshToken"`
	AccessExpiresIn  string  `json:"accessExpiresIn"`
	RefreshExpiresIn *string `json:"refreshExpiresIn"`
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (h *AuthService) createJwt(user *schema.User, key string, duration time.Duration) (string, time.Time, error) {
	signingKey := []byte(key)

	now := time.Now()
	expiresAt := now.Add(time.Duration(time.Hour * 24 * duration))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, utils.JwtClaims{
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

func (h *AuthService) createJwtPair(user *schema.User) (JwtPair, error) {
	accessToken, expAccess, accessErr := h.createJwt(user, config.C.Jwt.Secret, time.Duration(config.C.Jwt.Expire))
	if accessErr != nil {
		return JwtPair{}, accessErr
	}
	refreshToken, expRefresh, refreshErr := h.createJwt(user, config.C.Jwt.RefreshSecret, time.Duration(config.C.Jwt.RefreshExpire))
	if refreshErr != nil {
		return JwtPair{}, refreshErr
	}

	expRefreshFormatted := expRefresh.Format(time.RFC3339)
	expAccessFormatted := expAccess.Format(time.RFC3339)

	return JwtPair{
		AccessToken:      accessToken,
		RefreshToken:     &refreshToken,
		AccessExpiresIn:  expAccessFormatted,
		RefreshExpiresIn: &expRefreshFormatted,
	}, nil
}

func (h *AuthService) Login(email, password string) (*JwtPair, *helper.ServiceError) {
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

func (h *AuthService) Register(authData *serializer.AuthDataRegister) (*JwtPair, *helper.ServiceError) {
	ctx, cancel := utils.CreateContextTimeout(15)

	_, err := h.userRepo.FindByEmail(ctx, authData.Email)
	if err == nil {
		return nil, helper.NewServiceError(err, errors.EntityAlreadyExists)
	}

	cancel()

	hashedPwd, err := crypto.HashPassword(authData.Password)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.CryptoError)
	}

	ctxCreate, cancelCreate := utils.CreateContextTimeout(15)
	defer cancelCreate()

	user, err := h.userRepo.CreateEntity(ctxCreate, &schema.User{
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
func (h *AuthService) Refresh(refreshToken string) (*JwtPair, *helper.ServiceError) {
	parsed, err := utils.ParseJwtRefreshToken(refreshToken)
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

	return &JwtPair{
		AccessToken:      accessToken,
		AccessExpiresIn:  expAccess.Format(time.RFC3339),
		RefreshToken:     nil,
		RefreshExpiresIn: nil,
	}, nil
}
