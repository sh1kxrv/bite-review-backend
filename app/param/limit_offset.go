package param

import (
	"github.com/gofiber/fiber/v2"
)

func GetLimitOffset(c *fiber.Ctx) (limit int, offset int) {
	limit = c.QueryInt("limit", 10)
	offset = c.QueryInt("offset", 0)

	return limit, offset
}
