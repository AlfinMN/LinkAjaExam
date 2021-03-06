package master

import (
	"TestLinkAja/masters/api/controller"
	"TestLinkAja/masters/api/repositories"
	"TestLinkAja/masters/api/usecases"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func InitData(e *echo.Echo, db *sql.DB) {
	accRepository := repositories.InitAcountRepository(db)
	accUsecase := usecases.InitAccountUsecase(accRepository)
	accController := controller.NewAccountController(accUsecase)

	controller.AccountHandler(e, accController)
}
