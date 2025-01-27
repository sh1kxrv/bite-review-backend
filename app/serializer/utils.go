package serializer

import (
	"bitereview/utils"
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

func GetJwtUserLocal(c *fiber.Ctx) (utils.JwtClaims, error) {
	localUser := c.Locals("user")
	if localUser == nil {
		return utils.JwtClaims{}, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	parsedLocalUser, ok := localUser.(utils.JwtClaims)
	if !ok {
		return utils.JwtClaims{}, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	return parsedLocalUser, nil
}

// TODO: Требуется рефактор метода
func GetJwtUserLocalWithParsedID(c *fiber.Ctx) (utils.JwtClaims, primitive.ObjectID, error) {
	localUser := c.Locals("user")
	if localUser == nil {
		return utils.JwtClaims{}, primitive.NilObjectID, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	parsedLocalUser, ok := localUser.(utils.JwtClaims)
	if !ok {
		return utils.JwtClaims{}, primitive.NilObjectID, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	parsedID, err := primitive.ObjectIDFromHex(parsedLocalUser.ID)
	if err != nil {
		return utils.JwtClaims{}, primitive.NilObjectID, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	return parsedLocalUser, parsedID, nil
}
