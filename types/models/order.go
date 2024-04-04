package models

import "time"

type Item struct {
	ProductID int64 `json:"product_id" db:"product_id" binding:"required"`
	Amount    int64 `json:"amount" db:"amount" binding:"required"`
}

type Order struct {
	DbStandard
	PaymentID int64 `json:"payment_id" binding:"required"`
	ProductID int64 `json:"product_id" binding:"required"`
	Amount    int64 `json:"amount" binding:"required"`
}

type Payment struct {
	DbStandard
	UserID     int64     `json:"user_id"`
	Price      int64     `json:"price" binding:"required"`
	Code       string    `json:"code" binding:"required"`
	PaidOff    int8      `json:"paid_off" binding:"required"`
	Channel    string    `json:"channel" binding:"required"`
	FailReason string    `json:"fail_reason" binding:"required"`
	ExpiredAt  time.Time `json:"expired_at" binding:"required"`
}

type PaymentWithOrder struct {
	DbStandard
	Price      int64     `json:"price" binding:"required"`
	Code       string    `json:"code" binding:"required"`
	PaidOff    int8      `json:"paid_off" binding:"required"`
	Channel    string    `json:"channel" binding:"required"`
	FailReason string    `json:"fail_reason" binding:"required"`
	ExpiredAt  time.Time `json:"expired_at" binding:"required"`
	Orders     []Order   `json:"orders" binding:"required"`
}

type PayCallbackRequest struct {
	Code string `json:"code" db:"code" binding:"required"`
}
