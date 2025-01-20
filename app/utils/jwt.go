package utils

import (
	"bitereview/app/config"
	"bitereview/app/enum"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtClaims struct {
	Role     enum.Role          `json:"role"`
	ID       string             `json:"id"`
	ParsedID primitive.ObjectID `json:"-"`
	jwt.RegisteredClaims
}

func JwtTokenParse(secret string, token string) (JwtClaims, error) {
	claims := JwtClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return claims, err
}

func ParseJwtToken(token string) (JwtClaims, error) {
	return JwtTokenParse(config.C.Jwt.Secret, token)
}

func ParseJwtRefreshToken(token string) (JwtClaims, error) {
	return JwtTokenParse(config.C.Jwt.RefreshSecret, token)
}
