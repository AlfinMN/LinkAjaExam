package controller

import (
	"TestLinkAja/masters/api/models"
	"TestLinkAja/masters/api/usecases"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountController struct {
	accountUsecases usecases.AccountUsecases
}

func NewAccountController(usecase usecases.AccountUsecases) *AccountController {
	return &AccountController{accountUsecases: usecase}
}
func AccountHandler(e *echo.Echo, service *AccountController) {

	e.GET("/account/:accNum", service.CheckSaldo)
	e.POST("/account/:from_account_number/transfer", service.Transfer)
}

func (a *AccountController) CheckSaldo(c echo.Context) error {
	accNum := c.Param("accNum")
	fmt.Println(accNum)
	acc, err := a.accountUsecases.CheckSaldo(accNum)

	if err != nil {
		log.Println(err)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, acc)

}

func (a *AccountController) Transfer(c echo.Context) error {

	transfer := &models.Transfer{}
	transfer.Sender = c.Param("from_account_number")

	err := c.Bind(transfer)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = a.accountUsecases.Transfer(transfer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, "Transfer Succes")

}
