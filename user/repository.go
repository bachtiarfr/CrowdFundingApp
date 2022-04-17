package user

import (
	"fmt"

	"gorm.io/gorm"
)

/*
* create repository interface
 */
type Repository interface {
	SaveUser(user User) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUserByID(ID int) (User, error)
	UpdateUser(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveUser(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repository) GetUserByEmail(email string) (User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		message := fmt.Sprintf("User with email %s not found", email)
		fmt.Println(message)
	}
	return user, nil
}

func (r *repository) GetUserByID(ID int) (User, error) {
	var user User
	err := r.db.Where("ID = ?", ID).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) UpdateUser(user User) (User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}