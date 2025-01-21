package repository

import (
	"MiniWallet/internal/domain"
	"MiniWallet/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

// FindByUsername fetches a user by username
func (r *userRepository) FindByUsername(customerXid string) (*models.User, error) {
	var user models.User
	err := r.db.Where("customer_xid = ?", customerXid).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Create saves a new user in the database
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}
