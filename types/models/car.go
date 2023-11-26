package models

type Car struct {
	DbStandard
	Plat    string `json:"plat" binding:"required"`
	CarType string `json:"car_type" binding:"required"`
}

type CarByPlatRequest struct {
	Plat string `json:"plat" binding:"required"`
}
