package handler

import (
	"bitereview/app/config"
	"bitereview/app/crypto"
	"bitereview/app/errors"
	"bitereview/app/helper"
	"bitereview/app/repository"
	"bitereview/app/schema"
	"bitereview/app/serializer"
	"bitereview/app/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthHandler struct {
	UserRepo *repository.UserRepository
}

type JwtPair struct {
	AccessToken      string  `json:"accessToken"`
	RefreshToken     *string `json:"refreshToken"`
	AccessExpiresIn  string  `json:"accessExpiresIn"`
	RefreshExpiresIn *string `json:"refreshExpiresIn"`
}

func NewAuthHandler(userRepo *repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		UserRepo: userRepo,
	}
}

func (h *AuthHandler) createJwt(user *schema.User, key string, duration time.Duration) (string, time.Time, error) {
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

func (h *AuthHandler) createJwtPair(user *schema.User) (JwtPair, error) {
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

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	v, err := serializer.GetSerializedAuthLoginData(c)

	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	user, err := h.UserRepo.FindByEmail(c.Context(), v.Email)
	if err != nil {
		return helper.SendError(c, err, errors.RepositoryError)
	}

	if user == nil {
		return helper.SendError(c, err, errors.Forbidden)
	}

	if !crypto.CheckPasswordHash(v.Password, user.Password) {
		return helper.SendError(c, err, errors.Forbidden)
	}

	pair, err := h.createJwtPair(user)
	if err != nil {
		return helper.SendError(c, err, errors.JwtPairGenerationError)
	}

	h.UserRepo.UpdateLastSeen(c.Context(), user.ID)

	return helper.SendSuccess(c, pair)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	v, err := serializer.GetSerializedAuthRegisterData(c)

	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	_, err = h.UserRepo.FindByEmail(c.Context(), v.Email)
	if err == nil {
		return helper.SendError(c, err, errors.EntityAlreadyExists)
	}

	hashedPwd, err := crypto.HashPassword(v.Password)
	if err != nil {
		return helper.SendError(c, err, errors.CryptoError)
	}

	user, err := h.UserRepo.CreateUser(c.Context(), &schema.User{
		Email:    v.Email,
		Password: hashedPwd,
		Role:     "user",
		LastSeen: time.Now(),
	})

	if err != nil {
		return helper.SendError(c, err, errors.MakeRepositoryError("User"))
	}

	pair, err := h.createJwtPair(user)
	if err != nil {
		return helper.SendError(c, err, errors.JwtPairGenerationError)
	}

	return helper.SendSuccess(c, pair)
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	v, err := serializer.GetSerializedAuthRefreshData(c)

	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	parsed, err := utils.ParseJwtRefreshToken(v.RefreshToken)
	if err != nil {
		return helper.SendError(c, err, errors.JwtRefreshTokenInvalid)
	}

	parsedId, err := primitive.ObjectIDFromHex(parsed.ID)
	if err != nil {
		return helper.SendError(c, err, errors.ParseIDError)
	}

	user, err := h.UserRepo.FindByID(c.Context(), parsedId)
	if err != nil {
		return helper.SendError(c, err, errors.EntityNotExists)
	}

	accessToken, expAccess, accessErr := h.createJwt(user, config.C.Jwt.Secret, time.Duration(config.C.Jwt.Expire))
	if accessErr != nil {
		return helper.SendError(c, err, errors.JwtPairGenerationError)
	}

	return helper.SendSuccess(c, JwtPair{
		AccessToken:      accessToken,
		AccessExpiresIn:  expAccess.Format(time.RFC3339),
		RefreshToken:     nil,
		RefreshExpiresIn: nil,
	})
}

func (h *AuthHandler) RegisterRoutes(g fiber.Router) {
	authRoute := g.Group("/auth")

	authRoute.Post("/login", h.Login)
	authRoute.Post("/register", h.Register)
	authRoute.Post("/refresh", h.Refresh)
}
