package domain

import "MiniWallet/internal/models"

type UserRepository interface {
	FindByUsername(customerXid string) (*models.User, error)
	Create(user *models.User) error
}
