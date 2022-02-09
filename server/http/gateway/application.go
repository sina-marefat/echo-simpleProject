package gateway

import (
	"github.com/labstack/echo/v4"
	appDelivery "myProject/app/performance/delivery"
	repo "myProject/app/performance/repository"
	appUseCase "myProject/app/performance/usecase"
	db2 "myProject/db"
)

func RegisterApplication(e *echo.Echo) {
	db := db2.GetDB()
	repository := repo.NewRepo(db)
	appUC := appUseCase.New(appUseCase.Options{
		//Transaction: db.Begin(),
		Repo: repository,
	})

	appDelivery.RegisterHandlers(e, appUC)
}
