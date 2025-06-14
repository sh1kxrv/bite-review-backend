package user

import (
	"shared/database/mongodb/entity"
	"shared/database/mongodb/repository"
	"shared/errors"
	"shared/utils"
	"shared/utils/helper"

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

func (h *UserService) GetUserByID(userId primitive.ObjectID) (*entity.User, *helper.ServiceError) {
	timeoutCtx, cancel := utils.CreateContextTimeout(15)
	defer cancel()

	user, err := h.userRepo.GetEntityByID(timeoutCtx, userId)
	if err != nil {
		return nil, helper.NewServiceError(err, errors.MakeRepositoryError("User"))
	}

	return user, nil
}
