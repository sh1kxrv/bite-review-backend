package middleware

import (
	"bitereview/errors"
	"bitereview/helper"
	"bitereview/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JwtAuthMiddleware(c *fiber.Ctx) error {
	header := string(c.Request().Header.Peek("Authorization"))
	if header == "" {
		return helper.SendError(c, nil, errors.Unauthorized)
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" || len(headerParts[1]) == 0 {
		return helper.SendError(c, nil, errors.Unauthorized)
	}

	user, err := utils.ParseJwtToken(headerParts[1])
	if err != nil {
		return helper.SendError(c, nil, errors.JwtPairVerificationError)
	}

	c.Locals("user", user)

	return c.Next()
}
