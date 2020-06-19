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
		//根据类型id拿到博客对应的type
		blog.Category, err = GetTypeById(blog.TypeId)
		if err != nil {
			return
		}
		//根据用户id拿到博客对应的用户
		blog.User, err = GetUserById(blog.UserId)
		if err != nil {
			return
		}
		blogs = append(blogs, &blog)
	}
	return

}
