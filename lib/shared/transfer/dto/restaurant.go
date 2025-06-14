package dto

type CreateRestaurantDTO struct {
	Name        string          `json:"name" validate:"required"`
	Description *string         `json:"description"`
	Address     string          `json:"address"`
	City        string          `json:"city"`
	Country     string          `json:"country"`
	Site        string          `json:"site"`
	KitchenType []string        `json:"kitchenType"`
	Images      []string        `json:"images"`
	Metadata    *map[string]any `json:"metadata"`
}
