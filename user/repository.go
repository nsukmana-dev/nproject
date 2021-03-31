package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	ValidateEmail(email string) bool
	FindByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) ValidateEmail(email string) bool {
	var user []User
	_ = r.db.Where("email = ?", email).Find(&user)

	if len(user) > 0 {
		return true
	}
	return false
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
