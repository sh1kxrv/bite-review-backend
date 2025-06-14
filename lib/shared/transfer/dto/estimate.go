package dto

type CreateEstimateDTO struct {
	Name        string `json:"name" validate:"required,min=2,max=50"`
	Description string `json:"description" validate:"required,min=2,max=50"`
	Value       int    `json:"value" validate:"required,min=1,max=100"`
}
