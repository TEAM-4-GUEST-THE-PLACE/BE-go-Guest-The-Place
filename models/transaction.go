package models

import (
	"time"
)

type Transaction struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Price  string `json:"price"`
	Status string `json:"status" `

	Updated_at time.Time
	Created_at time.Time
}

type CreateTransactionRequest struct {
	ID     uint   `json:"id" gorm:"primaryKey"  `
	Price  string `json:"price" gorm:"price" `
	Status string `json:"status" gorm:"status" `

	Updated_ct time.Time
	Created_ct time.Time
}
