package handler

import (
	"bitereview/app/errors"
	"bitereview/app/helper"
	"bitereview/app/middleware"
	"bitereview/app/repository"
	"bitereview/app/serializer"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	UserRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepo: userRepo,
	}
}

func (h *UserHandler) GetMeHandler(c *fiber.Ctx) error {
	jwtClaims, err := serializer.GetJwtUserLocal(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	parsedId, err := primitive.ObjectIDFromHex(jwtClaims.ID)
	if err != nil {
		return helper.SendError(c, err, errors.ValidationError)
	}

	user, err := h.UserRepo.FindByID(c.Context(), parsedId)
	if err != nil {
		return helper.SendError(c, err, errors.MakeRepositoryError("User"))
	}
	return helper.SendSuccess(c, user)
}

func (h *UserHandler) RegisterRoutes(g fiber.Router) {
	userRoute := g.Group("/user", middleware.JwtAuthMiddleware)

	userRoute.Get("/me", h.GetMeHandler)
}
