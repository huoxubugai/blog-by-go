package controllers

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func IndexGet(c *gin.Context){
	typeList, err := service.GetAllType()
	if err!=nil{
		c.HTML(http.StatusInternalServerError,"error.html",nil)
		return
	}
	c.HTML(http.StatusOK,"index.html",gin.H{
		"types":typeList,
	})

}

var DB *gorm.DB