package models

type Product struct {
	DbStandard
	Name  string `json:"name" binding:"required"`
	Code  string `json:"code" binding:"required"`
	Price int64  `json:"price" binding:"required"`
}
