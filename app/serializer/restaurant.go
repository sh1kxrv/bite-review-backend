package serializer

import "github.com/gofiber/fiber/v2"

type CreateRestaurantDTO struct {
	Name        string          `json:"name" validate:"required"`
	Description string          `json:"description"`
	Address     string          `json:"address"`
	Location    string          `json:"location"`
	Country     string          `json:"country"`
	Site        string          `json:"site"`
	Metadata    *map[string]any `json:"metadata"`
}

func GetSerializedCreateRestaurant(c *fiber.Ctx) (CreateRestaurantDTO, error) {
	return GetSerializedBodyData[CreateRestaurantDTO](c)
}
