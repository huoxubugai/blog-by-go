package service

import (
	"blog/dao/db"
	"blog/models"
	"log"
)

func GetAllBlog() (blogs []*models.Blog, err error) {
	blogs, err = db.GetAllBlog()
	if err != nil {
		log.Fatal(err)
	}
	return
}
