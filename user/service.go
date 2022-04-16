package user

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(user UserInputRegister) (User, error)
	LoginUser(user UserInputLogin) (User, error)
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

func (s *service) LoginUser(input UserInputLogin) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.GetUserByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, fmt.Errorf("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil

}