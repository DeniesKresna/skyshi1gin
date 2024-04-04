package models

import "time"

type Item struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Amount    int64 `json:"amount" binding:"required"`
}

type Order struct {
	DbStandard
	PaymentID string `json:"payment_id" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
	Amount    int64  `json:"amount" binding:"required"`
}

type Payment struct {
	DbStandard
	Amount     int64      `json:"amount" binding:"required"`
	Price      int64      `json:"price" binding:"required"`
	Code       string     `json:"code" binding:"required"`
	PaidOff    int8       `json:"paid_off" binding:"required"`
	Channel    string     `json:"channel" binding:"required"`
	FailReason string     `json:"fail_reason" binding:"required"`
	ExpiredAt  *time.Time `json:"expired_at" binding:"required"`
}
