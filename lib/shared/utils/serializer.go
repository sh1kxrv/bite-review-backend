package utils

import (
	"bytes"
	"shared/transfer/dto"
	"shared/validator"

	"github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
)

func GetSerializedBodyData[T any](c *fiber.Ctx) (T, error) {
	var data T

	rawBody := c.Body()

	decoder := json.NewDecoder(bytes.NewReader(rawBody))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&data); err != nil {
		return data, err
	}

	validator := validator.GetValidatorInstance()
	if err := validator.Struct(data); err != nil {
		return data, err
	}

	return data, nil
}

func GetJwtUserLocal(c *fiber.Ctx) (dto.JwtClaims, error) {
	localUser := c.Locals("user")
	if localUser == nil {
		return dto.JwtClaims{}, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	parsedLocalUser, ok := localUser.(dto.JwtClaims)
	if !ok {
		return dto.JwtClaims{}, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	return parsedLocalUser, nil
}

// TODO: Требуется рефактор метода
func GetJwtUserLocalWithParsedID(c *fiber.Ctx) (dto.JwtClaims, primitive.ObjectID, error) {
	localUser := c.Locals("user")
	if localUser == nil {
		return dto.JwtClaims{}, primitive.NilObjectID, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	parsedLocalUser, ok := localUser.(dto.JwtClaims)
	if !ok {
		return dto.JwtClaims{}, primitive.NilObjectID, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	parsedID, err := primitive.ObjectIDFromHex(parsedLocalUser.ID)
	if err != nil {
		return dto.JwtClaims{}, primitive.NilObjectID, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	return parsedLocalUser, parsedID, nil
}
