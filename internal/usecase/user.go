package usecase

import (
	"MiniWallet/internal/domain"
	"MiniWallet/internal/models"
	"MiniWallet/pkg/utils"
	"errors"
)

type UserUsecase interface {
	Init(user *models.User) (string, error)
}

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) Init(user *models.User) (string, error) {
	// Check if the username already exists
	existingUser, _ := u.userRepo.FindByUsername(user.CustomerXid)
	if existingUser != nil {
		return "", errors.New("already enabled")
	}

	// Save the user
	err := u.userRepo.Create(user)
	if err != nil {
		return "", err
	}
	token, err := utils.GenerateToken(user.CustomerXid)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
