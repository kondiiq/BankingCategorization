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

type Currency struct {
	ID         string `json:"ID" binding:"required, max=3"`
	Name       string `json:"name" binding:"required"`
	Country    string `json:"country" binding:"required"`
	ContryCode string `json:"contryCode" binding:"max=3"`
}

//var receipts = []receipt{

//}
