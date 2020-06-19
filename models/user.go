package models

import "time"

type User struct {
	Id        int64     `gorm:"primary_key;auto_increment"`
	Nickname  string    `gorm:"not null;type:varchar(255)"`
	Username  string    `gorm:"not null;type:varchar(255)"`
	Password  string    `gorm:"not null;type:varchar(255)"`
	Email     string    `gorm:"type:varchar(30)"`
	Avatar    string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"not null;"`
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"`
}
