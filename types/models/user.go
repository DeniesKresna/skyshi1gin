package models

type User struct {
	DbStandard
	Email    string `gorm:"<-:create" json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   int64  `gorm:"column:role_id" json:"role_id" binding:"required"`
}

type Role struct {
	DbStandard
	Name string `json:"name"`
}

type EmailRequest struct {
	Email string `json:"email" binding:"required"`
}

type UserRole struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	RoleID   int64  `json:"role_id"`
	RoleName string `json:"role_name"`
}
