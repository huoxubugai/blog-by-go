package controllers

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BlogGet(c *gin.Context) {
	id := c.Param("id")
	pid, err := strconv.Atoi(id)
	checkError(c, err)
	blog, err := service.GetBlogById(pid)
	checkError(c, err)
	c.HTML(http.StatusOK, "blog.html", gin.H{
		"blog": blog,
	})
}
