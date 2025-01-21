package models

import "time"

type Wallet struct {
	ID         string     `json:"id" gorm:"primaryKey"`
	Balance    float64    `json:"balance"`
	Enabled    bool       `json:"enabled"`
	OwnedBy    string     `json:"owned_by"`
	DisabledAt *time.Time `json:"disabled_at, omitempty"`
	Status     string     `json:"status"`
}
