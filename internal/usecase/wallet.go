package usecase

import (
	"MiniWallet/internal/models"
	"MiniWallet/internal/repository"
	"MiniWallet/pkg/utils"
	"errors"
	"time"
)

type WalletUsecase interface {
	CreateWallet(customerXid string) (*models.Wallet, error)
	GetWalletAmount(customerXid string) (*models.Wallet, error)
	UpdateAmount(customerXid string, amount float64) (*models.Wallet, error)
	UpdateEnableStatus(customerXid string, status bool) (*models.Wallet, error)
}

type walletUsecase struct {
	walletRepo repository.WalletRepository
}

func NewWalletUsecase(walletRepo repository.WalletRepository) WalletUsecase {
	return &walletUsecase{walletRepo: walletRepo}
}

func (u *walletUsecase) CreateWallet(customerXid string) (*models.Wallet, error) {
	wallet := models.Wallet{
		ID:      utils.GenerateUUID(),
		Balance: 0,
		Enabled: true,
		OwnedBy: customerXid,
		Status:  "enabled",
	}

	return u.walletRepo.CreateWallet(wallet)
}

func (u *walletUsecase) GetWalletAmount(customerXid string) (*models.Wallet, error) {
	w, e := u.walletRepo.FindWalletByCustomerXid(customerXid)
	if e != nil {
		return nil, e
	}
	if w == nil {
		return nil, errors.New("wallet not found")
	}
	if !w.Enabled {
		return nil, errors.New("wallet is disabled")
	}
	return w, nil
}

func (u *walletUsecase) UpdateAmount(customerXid string, amount float64) (*models.Wallet, error) {
	wallet, err := u.walletRepo.FindWalletByCustomerXid(customerXid)
	if err != nil {
		return nil, err
	}

	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	wallet.Balance += amount
	return u.walletRepo.UpdateWallet(*wallet)
}

func (u *walletUsecase) UpdateEnableStatus(customerXid string, status bool) (*models.Wallet, error) {
	wallet, err := u.walletRepo.FindWalletByCustomerXid(customerXid)
	if err != nil {
		return nil, err
	}

	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	if wallet.Enabled == status {
		return nil, errors.New("wallet already in the same status")
	}

	if status == false {
		wallet.Status = "disabled"
		timeN := time.Now()
		wallet.DisabledAt = &timeN
	}

	wallet.Enabled = status
	return u.walletRepo.UpdateWallet(*wallet)
}
