package models

type TransactionRequest struct {
	Amount      float64 `json:"amount"`
	ReferenceId string  `json:"reference_id"`
}
