package db

import (
	"blog/config"
	"blog/models"
	"github.com/jinzhu/gorm"
)
import _ "github.com/go-sql-driver/mysql"

func InitDb(dns string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dns)
	db.SingularTable(true)
	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Blog{}, &models.Tag{})
	//table := db.HasTable("user")
	//if table==true{
	//	fmt.Println("存在user表")
	//}
	config.DB = db
	if err != nil {
		panic("连接数据库失败")
	}
	//defer db.Close()
	return db, nil
}
