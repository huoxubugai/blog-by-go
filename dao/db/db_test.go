package db

import (
	"blog/config"
	"blog/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

//var Db *gorm.DB
func init() {
	// parseTime=true 将mysql中时间类型，自动解析为go结构体中的时间类型
	// 不加报错
	dns := "root:wang6490918@tcp(localhost:3306)/blog?parseTime=true"
	db, err := gorm.Open("mysql", dns)
	db.SingularTable(true)
	config.DB = db
	if err != nil {
		panic(err)
	}

}

//分类测试
func TestGetAllType(t *testing.T) {
	types, err := GetAllType()
	if err != nil {
		panic(err)
	}
	for _, v := range types {
		t.Logf("id:%d type:%#v\n", v.Id, v)
	}
}

func TestGetTypeById(t *testing.T) {
	category, err := GetTypeById(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(category)
}

func TestInsertType(t *testing.T) {
	id, err := InsertType(&models.Category{Id: 4, Name: "ceshiceshi"})
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}

//博客数据库操作测试
func TestGetAllBlog(t *testing.T) {
	blogs, err := GetAllBlog()
	if err != nil {
		panic(err)
	}
	for _, blog := range blogs {
		t.Logf("title:%v,description:%v", blog.Title, blog.Description)
	}
}
