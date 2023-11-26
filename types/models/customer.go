package models

type Customer struct {
	DbStandard
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email"`
	Phone   string `json:"phone" binding:"required"`
	Address string `json:"address"`
}

type NameRequest struct {
	Name string `json:"name" binding:"required"`
}
