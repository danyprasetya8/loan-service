package user

import (
	"loan-service/internal/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(en *entity.User) error
	Save(en *entity.User) error
	Get(email string) *entity.User
}

type User struct {
	db *gorm.DB
}

func New(db *gorm.DB) IUserRepository {
	return &User{db}
}

func (u *User) Create(en *entity.User) error {
	return u.db.Create(en).Error
}

func (u *User) Save(en *entity.User) error {
	return u.db.Save(en).Error
}

func (u *User) Get(email string) (us *entity.User) {
	us = &entity.User{}

	err := u.db.Where("email = ?", email).
		First(us).
		Error

	if err != nil {
		return nil
	}

	return
}
