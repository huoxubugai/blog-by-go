package db

import (
	"blog/config"
	"blog/models"
	"log"
	"strconv"
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

func GetBlogById(id int) (models.Blog, error) {
	var blog models.Blog
	err := config.DB.Where("id=?", id).First(&blog).Error
	if err != nil {
		return models.Blog{}, err
	}
	blog.Category, err = GetTypeById(blog.TypeId)
	if err != nil {
		return models.Blog{}, err
	}
	blog.User, err = GetUserById(blog.UserId)
	if err != nil {
		return models.Blog{}, err
	}
	return blog, nil
}

func GetBlogByTypeId(id int) ([]models.Blog, error) {
	var blogs []models.Blog
	err := config.DB.Where("type_id=?", id).Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	for i, blog := range blogs {
		blog.Category, err = GetTypeById(blog.TypeId)
		if err != nil {
			return []models.Blog{}, err
		}
		blog.User, err = GetUserById(blog.UserId)
		if err != nil {
			return []models.Blog{}, err
		}
		blogs[i] = blog
	}
	return blogs, nil
}

func GetBlogByTagId(id int) ([]models.Blog, error) {
	var blogs []models.Blog
	strId := strconv.Itoa(id)
	param := "%" + strId + "%"
	rows, err := config.DB.Raw("select * from blog where tag_ids like ?", param).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var blog models.Blog
		config.DB.ScanRows(rows, &blog)
		blog.User, err = GetUserById(blog.UserId)
		if err != nil {
			return nil, err
		}
		blog.Category, err = GetTypeById(blog.TypeId)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}
