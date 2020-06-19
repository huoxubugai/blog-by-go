package models

import "time"

type User struct {
	Id        int64     `gorm:"primary_key;auto_increment"`
	Name      string    `gorm:"not null;type:varchar(255);"`
	Age       int       `gorm:"not null;type:int(3);default:0"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
}
