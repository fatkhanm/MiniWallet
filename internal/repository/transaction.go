package repository

import (
	"MiniWallet/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	AddToTransaction(transaction models.Transaction) error
	FindAllTransactionByCustomerXid(customerXid string) ([]models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) AddToTransaction(transaction models.Transaction) error {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) FindAllTransactionByCustomerXid(customerXid string) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Where("deposited_by = ? OR withdrawn_by = ?", customerXid, customerXid).
		Order("created_date").Find(&transaction).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return transaction, nil
}
