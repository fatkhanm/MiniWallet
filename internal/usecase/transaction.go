package usecase

import (
	"MiniWallet/internal/models"
	"MiniWallet/internal/repository"
	"MiniWallet/pkg/utils"
	"errors"
	"time"
)

type TransactionUsecase interface {
	ViewAllTransaction(customerXid string) ([]models.Transaction, error)
	Deposit(customerXid, referenceId string, amount float64) (*models.Transaction, error)
	Withdraw(customerXid, referenceId string, amount float64) (*models.Transaction, error)
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepository
	walletRepo      repository.WalletRepository
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepository, walletRepository repository.WalletRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: transactionRepo,
		walletRepo:      walletRepository,
	}
}

func (u *transactionUsecase) Deposit(customerXid, referenceId string, amount float64) (*models.Transaction, error) {
	// Get wallet by customerXid
	timeNow := time.Now()
	transaction := models.Transaction{
		ID:          utils.GenerateUUID(),
		ReferenceId: referenceId,
		Amount:      amount,
		DepositedBy: &customerXid,
		DepositedAt: &timeNow,
		CreatedDate: timeNow,
	}
	wallet, err := u.walletRepo.FindWalletByCustomerXid(customerXid)
	status := "success"
	if err != nil {
		status = "failed"

	}

	if wallet == nil {
		status = "failed"
	}
	walletId := ""
	if status != "failed" {
		// Add amount to wallet
		if !wallet.Enabled {
			return nil, errors.New("wallet is disabled")
		}
		wallet.Balance += amount
		walletId = wallet.ID
		if _, err := u.walletRepo.UpdateWallet(*wallet); err != nil {
			status = "failed"
		}
	}
	transaction.Status = status
	transaction.WalletId = walletId
	errTrans := u.transactionRepo.AddToTransaction(transaction)
	if errTrans != nil {
		return nil, errTrans
	}
	return &transaction, nil
}

func (u *transactionUsecase) Withdraw(customerXid, referenceId string, amount float64) (*models.Transaction, error) {
	timeNow := time.Now()
	transaction := models.Transaction{
		ID:          utils.GenerateUUID(),
		ReferenceId: referenceId,
		Amount:      amount,
		WithdrawnBy: &customerXid,
		WithdrawnAt: &timeNow,
		CreatedDate: timeNow,
		WalletId:    customerXid,
	}
	wallet, err := u.walletRepo.FindWalletByCustomerXid(customerXid)
	status := "success"
	if err != nil {
		status = "failed"
	}

	if wallet == nil {
		status = "failed"
	}
	walletId := ""
	if status != "failed" {
		// Add amount to wallet
		if !wallet.Enabled {
			return nil, errors.New("wallet is disabled")
		}
		wallet.Balance -= amount
		walletId = wallet.ID
		if _, err := u.walletRepo.UpdateWallet(*wallet); err != nil {
			status = "failed"
		}
	}
	transaction.Status = status
	transaction.WalletId = walletId
	errTrans := u.transactionRepo.AddToTransaction(transaction)
	if errTrans != nil {
		return nil, errTrans
	}
	return &transaction, nil
}

func (u *transactionUsecase) ViewAllTransaction(customerXid string) ([]models.Transaction, error) {
	w, e := u.walletRepo.FindWalletByCustomerXid(customerXid)
	if e != nil {
		return nil, e
	}
	if !w.Enabled {
		return nil, errors.New("wallet is disabled")
	}
	return u.transactionRepo.FindAllTransactionByCustomerXid(customerXid)
}
