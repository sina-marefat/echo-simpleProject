package models

import "time"

type Otp struct {
	UserId    int `gorm:"unique"`
	Code      int
	ExpiresAt time.Time
}

func (ds *Otp) TableName() string {
	return "otps"
}
