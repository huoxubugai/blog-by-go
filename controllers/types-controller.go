package controllers

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TypesGet(c *gin.Context) {
	id := c.Param("id")
	pid, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	//从菜单栏点进分类时，id=-1
	if pid == -1 {
		typeId, err := service.GetMaxBlogNumsInType()
		if err != nil {
			pid = 0
		} else {
			pid = typeId
		}
	}
	blogs, err := service.GetBlogByTypeId(pid)
	checkError(c, err)
	types, err := service.GetAllType()
	checkError(c, err)
	c.HTML(http.StatusOK, "types.html", gin.H{
		"blogs":         blogs,
		"types":         types,
		"currentTypeId": pid,
	})

}
