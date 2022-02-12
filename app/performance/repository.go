package performance

import (
	"github.com/labstack/echo/v4"
	"myProject/models"
)

type Repository interface {
	SignUp(ctx echo.Context, user *models.User) error
	Confirm(context echo.Context, id int, code int) error
	NewOtp(context echo.Context, otp models.Otp) error
	GetUserById(context echo.Context, id int) (models.User, error)
	CheckRenewOtp(ctx echo.Context, id int) (bool, error)
	GetOtpById(context echo.Context, id int) (models.Otp, error)
	GetUserByUsername(context echo.Context, username string) (models.User, error)
}
