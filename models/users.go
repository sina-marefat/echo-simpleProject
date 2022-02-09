package models

type User struct {
	ID       int    `gorm:"autoIncrement"`
	Username string `gorm:"unique"`
	Password string
	IsActive bool
}

func (ds *User) TableName() string {
	return "users"
}
