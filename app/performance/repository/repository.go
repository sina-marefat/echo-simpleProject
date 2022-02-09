package repository

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"myProject/app/performance"
	"myProject/models"
	"time"
)

type repo struct {
	db *gorm.DB
}

func (store *repo) SignUp(ctx echo.Context, user *models.User) error {
	err := store.db.Create(&user).Error
	return err
}

func NewRepo(db *gorm.DB) performance.Repository {
	return &repo{db}
}

func (store *repo) dbWithTimeOut(ctx context.Context, time time.Duration) *gorm.DB {
	timeoutContext, _ := context.WithTimeout(ctx, time)
	tx := store.db.WithContext(timeoutContext)
	return tx
}
