package user

import "gorm.io/gorm"

/*
* create repository interface
 */
type Repository interface {
	SaveUser(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository( db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveUser(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}