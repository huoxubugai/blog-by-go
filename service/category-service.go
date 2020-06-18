package service

import (
	"blog/dao/db"
	"blog/models"
)

func GetAllType()(typeList []*models.Category,err error){
	typeList, err = db.GetAllType()
	if err!=nil{
		return
	}
	return
}
