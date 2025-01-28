package utils

import (
	"bitereview/validator"
	"bytes"

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

func GetJwtUserLocal(c *fiber.Ctx) (JwtClaims, error) {
	localUser := c.Locals("user")
	if localUser == nil {
		return JwtClaims{}, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	parsedLocalUser, ok := localUser.(JwtClaims)
	if !ok {
		return JwtClaims{}, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	return parsedLocalUser, nil
}

// TODO: Требуется рефактор метода
func GetJwtUserLocalWithParsedID(c *fiber.Ctx) (JwtClaims, primitive.ObjectID, error) {
	localUser := c.Locals("user")
	if localUser == nil {
		return JwtClaims{}, primitive.NilObjectID, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	parsedLocalUser, ok := localUser.(JwtClaims)
	if !ok {
		return JwtClaims{}, primitive.NilObjectID, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	parsedID, err := primitive.ObjectIDFromHex(parsedLocalUser.ID)
	if err != nil {
		return JwtClaims{}, primitive.NilObjectID, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	return parsedLocalUser, parsedID, nil
}
