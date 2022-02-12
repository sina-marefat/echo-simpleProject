package delivery

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myProject/app/performance"
	"myProject/models"
	"myProject/response"
	"net/http"
	"strconv"
)

type handler struct {
	uc performance.UseCase
}

func RegisterHandlers(e *echo.Echo, uc performance.UseCase) {
	// create groups
	h := handler{uc}
	v1 := e.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", h.SignUp)
			auth.POST("/confirm", h.ConfirmOTP)
			auth.POST("/renew", h.RenewOtp)
		}
	}

}

func (h *handler) SignUp(ctx echo.Context) error {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.CustomApiResponse{
			"error": err.Error(),
		})
	}
	user.IsActive = false

	otp, err := h.uc.SignUp(ctx, user)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.CustomApiResponse{
			"error": err.Error(),
		})
	}
	id := otp.UserId
	print(id)
	verifyLink := fmt.Sprintf("http://localhost:8000/v1/auth/confirm?id=%d&code=%d", id, otp.Code)

	return ctx.JSON(http.StatusOK, response.CustomApiResponse{"data": "you have been signedUp",
		"code": otp.Code, "verification Link": verifyLink})

}

func (h *handler) VerifyAccount(context echo.Context) {

}

func (h *handler) ConfirmOTP(context echo.Context) error {
	id := context.QueryParam("id")
	code := context.QueryParam("code")
	if len(id) == 0 || len(code) == 0 {
		return context.JSON(http.StatusBadRequest, response.CustomApiResponse{"data": "mismatch in parameters"})
	}
	numId, _ := strconv.Atoi(id)
	numCode, _ := strconv.Atoi(code)
	err := h.uc.ConfirmOTP(context, numId, numCode)

	if err != nil {
		return context.JSON(http.StatusBadRequest, response.CustomApiResponse{
			"error": err.Error(),
		})
	}
	return context.JSON(http.StatusOK, response.CustomApiResponse{
		"message": "account Activated successfully",
	})

}

func (h *handler) RenewOtp(context echo.Context) error {
	id := context.QueryParam("id")
	numId, _ := strconv.Atoi(id)
	otp, err := h.uc.RenewOtp(context, numId)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response.CustomApiResponse{
			"error": err.Error(),
		})
	}
	verifyLink := fmt.Sprintf("http://localhost:8000/v1/auth/confirm?id=%d&code=%d", numId, otp.Code)
	return context.JSON(http.StatusOK, response.CustomApiResponse{"data": "you have been signedUp",
		"code": otp.Code, "verification Link": verifyLink})
}
