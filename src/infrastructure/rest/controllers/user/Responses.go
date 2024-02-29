package user

import "time"

type MessageResponse struct {
	Message string `json:"message"`
}

type ResponseUser struct {
	ID        int       `json:"id" example:"999"`
	UserName  string    `json:"user" example:"kuk"`
	Email     string    `json:"email" example:"kuk@mail.com"`
	FirstName string    `json:"firstName" example:"kuk"`
	LastName  string    `json:"lastName" example:"Mama"`
	Status    bool      `json:"status" example:"false"`
	CreatedAt time.Time `json:"createdAt,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
}
