package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AboutMe(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}
