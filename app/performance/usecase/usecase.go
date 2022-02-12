package usecase

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"math/rand"
	"myProject/app/performance"
	"myProject/models"
	"time"
)

type useCase struct {
	tx   *gorm.DB
	repo performance.Repository
}

func (u useCase) ConfirmOTP(context echo.Context, id int, code int) error {
	err := u.repo.Confirm(context, id, code)
	return err
}

func (u useCase) SignUp(ctx echo.Context, user models.User) (models.Otp, error) {
	err := u.repo.SignUp(ctx, &user)
	user, _ = u.repo.GetUserByUsername(ctx, user.Username)
	var otp models.Otp
	if err != nil {
		return otp, err
	}
	otp = NewOtp(user.ID)
	err = u.repo.NewOtp(ctx, otp)
	if err != nil {
		return otp, err
	}
	return otp, err
}

func (u useCase) RenewOtp(ctx echo.Context, id int) (models.Otp, error) {
	var otp models.Otp
	check, err := u.repo.CheckRenewOtp(ctx, id)
	if err != nil {
		return otp, err
	}
	if !check {
		otp, err = u.repo.GetOtpById(ctx, id)
		return otp, err
	}
	otp = NewOtp(id)
	err = u.repo.NewOtp(ctx, otp)
	return otp, err
}

func NewOtp(id int) models.Otp {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(9999-1000) + 1000
	otp := models.Otp{
		UserId:    id,
		Code:      code,
		ExpiresAt: time.Now().Add(time.Minute * 1),
	}
	return otp
}

type Options struct {
	//Transaction *gorm.DB
	Repo performance.Repository
}

func New(opts Options) performance.UseCase {
	return &useCase{
		repo: opts.Repo,
	}
}
