package models

import (
	"time"
)

type Execution struct {
	DbStandard
	ExecuterID  int64      `json:"executer_id"`
	CarID       int64      `json:"car_id" binding:"required"`
	CustomerID  int64      `json:"customer_id" binding:"required"`
	Description string     `json:"description" binding:"required"`
	Price       int64      `json:"price"`
	Paid        int64      `json:"paid"`
	PaidOff     int        `json:"paid_off"`
	PaidAt      *time.Time `json:"paid_at"`
	ExecutedAt  *time.Time `json:"executed_at"`
	FinishAt    *time.Time `json:"finish_at"`
	UpdatedBy   string     `json:"updated_by"`
}

type ExecutionFilter struct {
	ExecuterName  string     `json:"executer_name"`
	CarPlat       string     `json:"car_plat"`
	CustomerName  string     `json:"customer_name"`
	CustomerPhone string     `json:"customer_phone"`
	CustomerEmail string     `json:"customer_email"`
	Price         []int64    `json:"price"`
	Paid          []int64    `json:"paid"`
	PaidOff       *bool      `json:"paid_off"`
	PaidFrom      *time.Time `json:"paid_from"`
	PaidTo        *time.Time `json:"paid_to"`
	ExecutedFrom  *time.Time `json:"executed_from"`
	ExecutedTo    *time.Time `json:"executed_to"`
	FinishFrom    *time.Time `json:"finish_from"`
	FinishTo      *time.Time `json:"finish_to"`
}

type ExecutionSearchResponse struct {
	ID            int64      `json:"id" xlsxField:"A:ID"`
	ExecuterID    int64      `json:"executer_id"`
	ExecuterName  string     `json:"executer_name" gorm:"column:user_name" xlsxField:"B:Nama Pekerja"`
	CarID         int64      `json:"car_id"`
	CarPlat       string     `json:"plat" gorm:"column:plat" xlsxField:"C:Plat Nomor"`
	CustomerID    int64      `json:"customer_id"`
	CustomerName  string     `json:"customer_name" xlsxField:"D:Nama Pelanggan"`
	CustomerPhone string     `json:"customer_phone" xlsxField:"E:Telpon Pelanggan"`
	CustomerEmail string     `json:"customer_email" xlsxField:"F:Email Pelanggan"`
	Description   string     `json:"description" xlsxField:"G:Keterangan"`
	Price         int64      `json:"price" xlsxField:"H:Harga"`
	Paid          int64      `json:"paid" xlsxField:"I:Terbayar"`
	PaidOff       *int       `json:"paid_off" xlsxField:"J:Lunas"`
	PaidAt        *time.Time `json:"paid_at" xlsxField:"K:Dibayar pada"`
	ExecutedAt    *time.Time `json:"executed_at" xlsxField:"L:Dikerjakan pada"`
	FinishAt      *time.Time `json:"finish_at" xlsxField:"M:Diselesaikan pada"`
}
