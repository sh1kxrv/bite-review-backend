package middleware

import (
	"shared/enum"
	"shared/errors"
	"shared/transfer/dto"
	"shared/utils/helper"
	"slices"

	"github.com/gofiber/fiber/v2"
)

func CreateRoleMiddleware(roles ...enum.Role) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(dto.JwtClaims)
		if !ok || !slices.Contains(roles, user.Role) {
			return helper.SendError(c, nil, errors.Forbidden)
		}
		return c.Next()
	}
}
