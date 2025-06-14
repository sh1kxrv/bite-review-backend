package dto

type ReviewDTO struct {
	Summary string `json:"summary" validate:"required,min=64,max=1024"`
}
