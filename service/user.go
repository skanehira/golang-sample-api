package service

import (
	"github.com/skanehira/golang-sample-api/model"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewServiceUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (u *User) Users() ([]model.UserInfo, error) {
	users := []model.UserInfo{}
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
