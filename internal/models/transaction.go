package models

import "time"

type Transaction struct {
	ID          string     `gorm:"primaryKey" json:"id"`
	DepositedBy *string    `json:"deposited_by,omitempty"`
	Status      string     `gorm:"not null" json:"status"`
	Amount      float64    `gorm:"not null" json:"amount" `
	ReferenceId string     `gorm:"not null" json:"reference_id"`
	WithdrawnBy *string    `json:"withdrawn_by,omitempty"`
	WalletId    string     `gorm:"not null" json:"wallet_id"`
	WithdrawnAt *time.Time `json:"withdrawn_at,omitempty"`
	DepositedAt *time.Time `json:"deposited_at,omitempty"`
	CreatedDate time.Time  `json:"created_date,omitempty"`
}
