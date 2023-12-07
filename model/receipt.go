package model

import (
	"gorm.io/gorm"
	"time"
)

type Receipt struct {
	gorm.Model
	ID              string    `json:"id" binding:"required"`
	Sum             float32   `json:"sum" binding:"required"`
	Address         string    `json:"address" binding:"required"`
	PaymentType     string    `json:"paymentType" binding:"required"`
	TransactionDate time.Time `json:"transaction_date"`
	Currency        string    `json:"currency" binding:"max=3" binding:"required"`
}

//var receipts = []receipt{

//}
