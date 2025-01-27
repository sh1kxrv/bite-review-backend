package service

import (
	"bitereview/errors"
	"bitereview/helper"
	"bitereview/repository"
	"bitereview/schema"
	"bitereview/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (h *UserService) GetUserByID(userId primitive.ObjectID) (*schema.User, *helper.ServiceError) {
	timeoutCtx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	user, err := h.userRepo.GetEntityByID(timeoutCtx, userId)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.MakeRepositoryError("User"))
	}

	return user, nil
}
