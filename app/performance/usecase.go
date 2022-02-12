package performance

import (
	"github.com/labstack/echo/v4"
	"myProject/models"
)

type UseCase interface {
	SignUp(ctx echo.Context, user models.User) (models.Otp, error)
	ConfirmOTP(context echo.Context, id int, code int) error
	RenewOtp(ctx echo.Context, id int) (models.Otp, error)
}
