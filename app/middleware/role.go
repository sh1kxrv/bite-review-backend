package middleware

import (
	"bitereview/app/enum"
	"bitereview/app/errors"
	"bitereview/app/helper"
	"bitereview/app/utils"
	"slices"

	"github.com/gofiber/fiber/v2"
)

func CreateRoleMiddleware(roles ...enum.Role) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(utils.JwtClaims)
		if !ok || !slices.Contains(roles, user.Role) {
			return helper.SendError(c, nil, errors.Forbidden)
		}
		return c.Next()
	}
}
