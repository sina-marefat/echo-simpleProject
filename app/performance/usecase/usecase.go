package usecase

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"myProject/app/performance"
	"myProject/models"
)

type useCase struct {
	tx   *gorm.DB
	repo performance.Repository
}

func (u useCase) SignUp(ctx echo.Context, user models.User) error {
	err := u.repo.SignUp(ctx, &user)
	if err != nil {
		return err
	}
	return nil
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
