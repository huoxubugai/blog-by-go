package service

import (
	"blog/dao/db"
	"blog/models"
)

func GetAllType() (typeList []*models.Category, err error) {
	typeList, err = db.GetAllType()
	if err != nil {
		return
	}
	return
}

func GetMaxBlogNumsInType() (int, error) {
	typeId, err := db.GetMaxBlogNumsInType()
	if err != nil {
		return 0, err
	}
	return typeId, nil
}
