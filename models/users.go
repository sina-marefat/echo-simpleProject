package models

import "time"

type User struct {
	ID          int    `gorm:"primaryKey"`
	Username    string `gorm:"unique"`
	Password    string
	IsActive    bool
	ActivatedAt time.Time
}

func (ds *User) TableName() string {
	return "users"
}
