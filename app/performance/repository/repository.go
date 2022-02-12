package repository

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"myProject/app/performance"
	"myProject/models"
	"time"
)

type repo struct {
	db *gorm.DB
}

func (store *repo) GetUserByUsername(context echo.Context, username string) (models.User, error) {
	var user models.User
	err := store.db.First(&user, "username = ?", username)
	if err != nil {
		return user, errors.New("no such username found")
	}
	return user, nil
}

func (store *repo) CheckRenewOtp(ctx echo.Context, id int) (bool, error) {
	var otp models.Otp
	err := store.db.First(&otp, id)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return false, errors.New("this action is invalid")
	}
	if otp.ExpiresAt.Before(time.Now()) {
		return false, nil
	}
	store.db.Delete(&otp)
	return true, nil
}

func (store *repo) GetOtpById(context echo.Context, id int) (models.Otp, error) {
	var otp models.Otp
	err := store.db.First(&otp, id)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return otp, errors.New("no otp found")
	}
	return otp, nil
}
func (store *repo) GetUserById(context echo.Context, id int) (models.User, error) {
	var user models.User
	err := store.db.First(&user, id)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("no users found")
	}
	return user, nil
}

func (store *repo) NewOtp(context echo.Context, otp models.Otp) error {
	err := store.db.Create(&otp).Error
	if err != nil {
		return err
	}
	return nil
}

func (store *repo) Confirm(context echo.Context, id int, code int) error {
	var otp models.Otp
	err := store.db.First(&otp, id)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return errors.New("no otp found try to renew or signup")
	}
	if otp.Code != code {
		return errors.New("wrong code please try again")
	}
	if otp.ExpiresAt.Before(time.Now()) {
		return errors.New("otp is expired goto : http://localhost:8000/v1/auth/renew")
	}
	err = store.db.Model(&models.User{}).Where("id = ?", id).Update("is_active", true).Update("activated_at", time.Now())
	store.db.Where("user_id = ?", id).Delete(&otp)
	return nil
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
