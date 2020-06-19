package db

import (
	"blog/config"
	"blog/models"
)

func GetUserById(id int) (models.User, error) {
	var user models.User
	err := config.DB.Where("id=?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
