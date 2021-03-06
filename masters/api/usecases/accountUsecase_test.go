package usecases

import (
	"TestLinkAja/masters/api/models"
	"TestLinkAja/masters/api/repositories"
	"TestLinkAja/masters/utils"
	"database/sql"
	"log"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccountUsecase struct {
	mock.Mock
}

var account = &models.Account{
	AccountNumber: "555002",
	CustomerName:  "Elon Musk",
	Balance:       30000,
}

var transfer = &models.Transfer{
	Receiver: "555001",
	Amount:   200,
	Sender:   "555002",
}

func NewDBMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error occurred while opening a mock database connection")
	}
	return db, mock
}

func TestCheckSaldoSucces(t *testing.T) {
	db, mock := NewDBMock()
	repository := &repositories.Repositories{db}

	defer repository.DB.Close()

	usecase := InitAccountUsecase(repository)

	rows := sqlmock.NewRows([]string{"account_number", "customer_name", "balance"}).
		AddRow(account.AccountNumber, account.CustomerName, account.Balance)

	mock.ExpectQuery(regexp.QuoteMeta(utils.CHECKSALDO)).
		WithArgs(account.AccountNumber).
		WillReturnRows(rows)

	accountData, err := usecase.CheckSaldo(account.AccountNumber)
	assert.NotNil(t, accountData)
	assert.NoError(t, err)
}

func TestCheckSaldoFailure(t *testing.T) {
	db, mock := NewDBMock()
	repository := &repositories.Repositories{db}

	defer repository.DB.Close()

	usecase := InitAccountUsecase(repository)

	rows := sqlmock.NewRows([]string{"account_number", "customer_name", "balance"})

	mock.ExpectQuery(regexp.QuoteMeta(utils.CHECKSALDO)).WithArgs(account.AccountNumber).
		WillReturnRows(rows)

	accountData, err := usecase.CheckSaldo(account.AccountNumber)
	assert.Nil(t, accountData)
	assert.Error(t, err)
}

func TestTransferSucces(t *testing.T) {
	db, mock := NewDBMock()
	repository := &repositories.Repositories{db}

	defer repository.DB.Close()

	amount := strconv.Itoa(transfer.Amount)
	balance := strconv.Itoa(account.Balance)

	usecase := &AccountUsecaseImpl{repository}
	rows := sqlmock.NewRows([]string{"account_number", "customer_number", "balance"}).
		AddRow(account.AccountNumber, account.CustomerName, balance)

	mock.ExpectQuery(regexp.QuoteMeta(utils.CHECKSALDO)).WithArgs(account.AccountNumber).WillReturnRows(rows)

	rowtansfer := sqlmock.NewRows([]string{"amount", "account"}).
		AddRow(amount, account.AccountNumber)

	mock.ExpectBegin()

	mock.ExpectQuery(regexp.QuoteMeta(utils.SENDTRANSFER)).
		WithArgs(transfer.Amount, transfer.Sender).
		WillReturnRows(rowtansfer)
	mock.ExpectRollback()

	mock.ExpectQuery(regexp.QuoteMeta(utils.RECEIVETRANSFER)).
		WithArgs(transfer.Amount, transfer.Receiver).
		WillReturnRows(rowtansfer)
	mock.ExpectRollback()

	mock.ExpectCommit()

	transaction := usecase.Transfer(transfer)
	println(transaction)
	// assert.NoError(t, transaction)
}
