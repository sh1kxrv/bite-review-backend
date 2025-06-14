package utils

import (
	"shared/transfer/dto"

	"github.com/golang-jwt/jwt/v5"
)

func ParseJwtToken(token string, secret string) (dto.JwtClaims, error) {
	claims := dto.JwtClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return claims, err
}
