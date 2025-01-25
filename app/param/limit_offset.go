package param

import (
	"github.com/gofiber/fiber/v2"
)

func GetLimitOffset(c *fiber.Ctx) (limit int64, offset int64) {
	limitInt := c.QueryInt("limit", 10)
	offsetInt := c.QueryInt("offset", 0)

	return int64(limitInt), int64(offsetInt)
}
