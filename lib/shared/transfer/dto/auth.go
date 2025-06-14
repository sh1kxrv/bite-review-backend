package dto

import (
	"shared/enum"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtClaims struct {
	Role     enum.Role          `json:"role"`
	ID       string             `json:"id"`
	ParsedID primitive.ObjectID `json:"-"`
	jwt.RegisteredClaims
}

type AuthDataLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthDataRegister struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string `json:"lastName" validate:"required,min=2,max=50"`
}

type AuthDataRefresh struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
