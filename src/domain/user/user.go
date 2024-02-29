package user

import "time"

type User struct {
	ID           int
	UserName     string
	Email        string
	FirstName    string
	LastName     string
	Status       bool
	Role         string
	HashPassword string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type NewUser struct {
	UserName  string
	Email     string
	FirstName string
	LastName  string
	Role      string
	Password  string
	Status    bool
}

func (newUser *NewUser) ToDomainMapper() *User {
	return &User{
		UserName:  newUser.UserName,
		Email:     newUser.Email,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Role:      newUser.Role,
	}
}

type Service interface {
	GetAll() (*[]User, error)
	GetByID(id int) (*User, error)
	Create(newUser *NewUser) (*User, error)
	GetOneByMap(userMap map[string]interface{}) (*User, error)
	Delete(id int) error
	Update(id int, userMap map[string]interface{}) (*User, error)
}
