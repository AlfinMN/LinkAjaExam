package controller

import (
	"TestLinkAja/masters/api/models"
	"TestLinkAja/masters/api/repositories"
	"TestLinkAja/masters/api/usecases"
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var account = &models.Account{
	AccountNumber: "555005",
	CustomerName:  "Elon Musk",
	Balance:       30000,
}
var Transfer = &models.Transfer{
	Receiver: "555006",
	Amount:   200,
	Sender:   "555005",
}

func NewDBMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error occurred while opening a mock database connection")
	}
	return db, mock
}

func TestCheckSaldo(t *testing.T) {
	db, _ := NewDBMock()
	repository := &repositories.Repositories{db}
	usecase := usecases.InitAccountUsecase(repository)

	defer repository.DB.Close()

	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	record := httptest.NewRecorder()

	c := e.NewContext(request, record)
	c.SetPath("/account/:id")
	c.SetParamNames("id")
	c.SetParamValues(account.AccountNumber)

	handle := NewAccountController(usecase)
	if assert.NoError(t, handle.CheckSaldo(c)) {
		assert.Equal(t, http.StatusInternalServerError, record.Code)
	}
}

func TestTransfer(t *testing.T) {
	db, _ := NewDBMock()
	repository := &repositories.Repositories{db}
	usecase := usecases.InitAccountUsecase(repository)

	defer repository.DB.Close()

	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	record := httptest.NewRecorder()

	c := e.NewContext(request, record)
	c.SetPath("/account/:id/transfer")
	c.SetParamNames("id")
	c.SetParamValues(account.AccountNumber)
	handle := NewAccountController(usecase)
	if assert.NoError(t, handle.Transfer(c)) {
		assert.Equal(t, http.StatusBadRequest, record.Code)
	}
}
