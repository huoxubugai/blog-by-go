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

func GetBlogById(id int) (models.Blog, error) {
	blog, err := db.GetBlogById(id)
	if err != nil {
		return models.Blog{}, err
	}
	return blog, nil
}

func GetBlogByTypeId(id int) ([]models.Blog, error) {
	blogs, err := db.GetBlogByTypeId(id)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}
