package utils

import (
	"bitereview/app/config"
	"bitereview/app/enum"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtClaims struct {
	Role     enum.Role          `json:"role"`
	ID       string             `json:"id"`
	ParsedID primitive.ObjectID `json:"-"`
	jwt.RegisteredClaims
}

func jwtTokenParse(secret, token string) (JwtClaims, error) {
	claims := JwtClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return claims, err
}

func jwtCheckValidation(secret, tokenRaw string) (bool, error) {
	token, _, err := jwt.NewParser().ParseUnverified(tokenRaw, jwt.MapClaims{})
	if err != nil {
		return false, fmt.Errorf("token parsing failed: %w", err)
	}

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return false, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	err = token.Method.Verify(token.Raw, token.Signature, secret)
	if err != nil {
		return false, fmt.Errorf("signature verification failed: %w", err)
	}
	return true, nil
}

func ValidateJwtToken(token string) (bool, error) {
	return jwtCheckValidation(config.C.Jwt.Secret, token)
}

func ValidateJwtRefreshToken(token string) (bool, error) {
	return jwtCheckValidation(config.C.Jwt.RefreshSecret, token)
}

func ParseJwtToken(token string) (JwtClaims, error) {
	return jwtTokenParse(config.C.Jwt.Secret, token)
}

func ParseJwtRefreshToken(token string) (JwtClaims, error) {
	return jwtTokenParse(config.C.Jwt.RefreshSecret, token)
}
