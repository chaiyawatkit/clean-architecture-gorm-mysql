package user

import (
	userDomain "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/user"
	userRepository "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/infrastructure/repository/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserRepository userRepository.Repository
}

var _ userDomain.Service = &Service{}

func (s *Service) GetAll() (*[]userDomain.User, error) {
	return s.UserRepository.GetAll()
}

func (s *Service) GetByID(id int) (*userDomain.User, error) {
	return s.UserRepository.GetByID(id)
}

func (s *Service) Create(newUser *userDomain.NewUser) (*userDomain.User, error) {
	domain := newUser.ToDomainMapper()

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return &userDomain.User{}, err
	}
	domain.HashPassword = string(hash)
	domain.Status = true

	return s.UserRepository.Create(domain)
}

func (s *Service) GetOneByMap(userMap map[string]any) (*userDomain.User, error) {
	return s.UserRepository.GetOneByMap(userMap)
}

func (s *Service) Delete(id int) error {
	return s.UserRepository.Delete(id)
}

func (s *Service) Update(id int, userMap map[string]any) (*userDomain.User, error) {
	return s.UserRepository.Update(id, userMap)
}
