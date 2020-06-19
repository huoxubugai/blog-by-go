package service

import (
	"blog/dao/db"
	"blog/models"
)

func GetAllTag() ([]models.Tag, error) {
	tags, err := db.GetAllTag()
	if err != nil {
		return nil, err
	}
	return tags, nil
}
