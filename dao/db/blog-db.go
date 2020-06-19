package db

import (
	"blog/config"
	"blog/models"
	"log"
)

func GetAllBlog() (blogs []*models.Blog, err error) {
	rows, err := config.DB.Raw("select * from blog").Rows()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var blog models.Blog
		config.DB.ScanRows(rows, &blog)
		blog.Category, err = GetTypeById(blog.TypeId)
		if err != nil {
			return
		}
		blogs = append(blogs, &blog)
	}
	return

}
