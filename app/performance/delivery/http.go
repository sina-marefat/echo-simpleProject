package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"myProject/app/performance"
	"myProject/models"
	"myProject/response"
	"net/http"
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
		}
	}

}

func (h *handler) SignUp(ctx echo.Context) error {
	span := opentracing.StartSpan("get")
	defer span.Finish()
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.CustomApiResponse{
			"error": err.Error(),
		})
	}
	user.IsActive = false

	err := h.uc.SignUp(ctx, user)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.CustomApiResponse{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.CustomApiResponse{"data": "you have been signedUp"})

}
