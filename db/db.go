package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"myProject/models"
)

var db *gorm.DB
var _ error

func Init() {
	dsn := "host=localhost user=postgres password=secret dbname=myapp port=54321 sslmode=disable TimeZone=Asia/Tehran"
	initDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// defer db.Close()
	db = initDB
	if err != nil {
		panic("DB Connection Error")
	}
	err = db.AutoMigrate(&models.User{}, &models.Otp{})

	if err != nil {
		print("failed to migrate")
	}

}

func GetDB() *gorm.DB {
	return db
}
