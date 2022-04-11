package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(user UserInputRegister) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input UserInputRegister) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	user.Role = "user"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	newUser, err := s.repository.SaveUser(user)
	if err != nil {
		return user, err
	}
	return newUser, nil

}