package repository

import (
	"MiniWallet/internal/models"

	"gorm.io/gorm"
)

type WalletRepository interface {
	FindWalletByCustomerXid(customerXid string) (*models.Wallet, error)
	UpdateWallet(wallet models.Wallet) (*models.Wallet, error)
	CreateWallet(wallet models.Wallet) (*models.Wallet, error)
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{db: db}
}

func (r *walletRepository) CreateWallet(wallet models.Wallet) (*models.Wallet, error) {
	err := r.db.Create(&wallet).Error
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (r *walletRepository) FindWalletByCustomerXid(customerXid string) (*models.Wallet, error) {
	var wallet models.Wallet
	err := r.db.First(&wallet, "owned_by = ?", customerXid).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &wallet, nil
}

func (r *walletRepository) UpdateWallet(wallet models.Wallet) (*models.Wallet, error) {
	err := r.db.Save(&wallet).Error
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}
