package models

type AuthResponse struct {
	Token     string `json:"token"`
	User      User   `json:"user"`
	Role      Role   `json:"role"`
	ExpiredAt string `json:"expired_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
