package auth

type LoginRequest struct {
	Email    string `json:"email" example:"test@gmail.com" gorm:"unique" binding:"required"`
	Password string `json:"password" example:"Password123" binding:"required"`
}

type AccessTokenRequest struct {
	RefreshToken string `json:"refreshToken" example:"badbunybabybebe" binding:"required"`
}
