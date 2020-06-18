package models

type Category struct {
	Id int `gorm:"column:id;type:int;primary_key;auto_increment"`
	Name string `gorm:"column:name;type:varchar(255)"`
}
