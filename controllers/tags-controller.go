package controllers

import (
	"blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TagsGet(c *gin.Context) {
	id := c.Param("id")
	pid, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	if pid == -1 {
		pid = 1
	}
	blogs, err := service.GetBlogByTagId(pid)
	checkError(c, err)
	tags, err := service.GetAllTag()
	checkError(c, err)

	c.HTML(http.StatusOK, "tags.html", gin.H{
		"blogs": blogs,
		"tags":  tags,
	})
}
