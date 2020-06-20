package models

import "time"

type Blog struct {
	Id           int       `gorm:"column:id;type:int;primary_key;auto_increment"`
	Title        string    `gorm:"column:title;type:varchar(255);not null;"`
	Content      string    `gorm:"column:content;type:longtext;"`
	FirstPicture string    `gorm:"column:first_picture;type:varchar(255);"`
	Description  string    `gorm:"column:description;type:varchar(255);"`
	Views        int       `gorm:"column:views;type:int"`
	TypeId       int       `gorm:"column:type_id;type:int"`
	UserId       int       `gorm:"column:user_id;type:int;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;not null"`
	UpdatedAt    time.Time `gorm:"column:updated_at;default:current_timestamp"`
	TagIds       string    `gorm:"column:tag_ids;type:varchar(30)"`

	Category `gorm:"-"`
	User     `gorm:"-"`
}
