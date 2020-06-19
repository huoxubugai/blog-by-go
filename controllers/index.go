package controllers

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexGet(c *gin.Context) {
	//获取分类
	typeList, err := service.GetAllType()
	checkError(c, err)
	//获取博客
	blogs, err := service.GetAllBlog()
	checkError(c, err)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"types": typeList,
		"blogs": blogs,
	})

}

func checkError(c *gin.Context, err error) {
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
}
