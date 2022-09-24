package service

import (
	"github.com/giovannitgl/video-services/user-service/internal/config"
	"github.com/giovannitgl/video-services/user-service/internal/entities"
)

func UserGetOne(id uint) *entities.User {
	var user entities.User
	err := config.DB.Db.First(&user, id).Error
	if err != nil {
		return nil
	}
	return &user
}

func UserGetLogin(email, password string) *entities.User {
	var user entities.User
	err := config.DB.Db.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return nil
	}
	return &user
}

func UserCreate(user *entities.User) error {
	return config.DB.Db.Create(user).Error
}
