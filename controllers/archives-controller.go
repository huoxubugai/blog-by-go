package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArchivesGet(c *gin.Context) {
	c.HTML(http.StatusOK, "archives.html", gin.H{})
}
