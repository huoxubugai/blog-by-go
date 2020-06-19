package models

type Tag struct {
	Id       uint   `gorm:"not null;primary_key;auto_increment;type:int"`
	Name     string `gorm:"not null;type:varchar(30);"`
	BlogNums uint   `gorm:"type:int"`
}
