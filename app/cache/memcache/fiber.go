package memcache

import "github.com/gofiber/fiber/v2"

func GetMemoryCache(c *fiber.Ctx) *MemoryCache {
	return c.Locals("memoryCache").(*MemoryCache)
}
