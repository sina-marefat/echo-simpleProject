package performance

import (
	"github.com/labstack/echo/v4"
	"myProject/models"
)

type UseCase interface {
	SignUp(ctx echo.Context, user models.User) error
}
