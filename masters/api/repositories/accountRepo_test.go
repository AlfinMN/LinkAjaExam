package repositories

import (
	"TestLinkAja/masters/api/models"
	"TestLinkAja/masters/utils"
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var account = &models.Account{
	AccountNumber: "555005",
	CustomerName:  "Elon Musk",
	Balance:       30000,
}
var transfer = &models.Transfer{
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

func TestNewAccountRepo(t *testing.T) {
	db, _ := NewDBMock()
	repository := InitAcountRepository(db)
	assert.NotNil(t, repository)
}

func TestCheckSaldoSucces(t *testing.T) {
	db, mock := NewDBMock()
	repository := &Repositories{db}
	defer repository.DB.Close()

	e := strconv.Itoa(account.Balance)

	rows := sqlmock.NewRows([]string{"account_number", "customer_name", "balance"}).AddRow(account.AccountNumber, account.CustomerName, e)

	mock.ExpectQuery(regexp.QuoteMeta(utils.CHECKSALDO)).WithArgs(account.AccountNumber).WillReturnRows(rows)

	accountData, err := repository.CheckSaldo(account.AccountNumber)
	assert.NotNil(t, accountData)
	assert.NoError(t, err)
}

func TestCheckSaldoFailure(t *testing.T) {
	db, mock := NewDBMock()
	repository := &Repositories{db}

	defer repository.DB.Close()

	mock.ExpectQuery(utils.CHECKSALDO).WithArgs(account.AccountNumber).WillReturnError(fmt.Errorf("there is an error"))
	_, err := repository.CheckSaldo(account.AccountNumber)

	if assert.NotNil(t, err) {
		assert.Equal(t, "there is an error", err.Error())
	}

}

func TestTransferSucces(t *testing.T) {
	db, mock := NewDBMock()
	repository := &Repositories{db}

	defer repository.DB.Close()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(utils.SENDTRANSFER)).WithArgs(transfer.Amount, transfer.Sender).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(utils.RECEIVETRANSFER)).WithArgs(transfer.Amount, transfer.Receiver).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repository.Transfer(transfer)
	assert.NoError(t, err)
}

func TestSendTransferFailure(t *testing.T) {
	db, mock := NewDBMock()
	repository := &Repositories{db}

	defer repository.DB.Close()
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(utils.SENDTRANSFER)).
		WithArgs(transfer.Amount, transfer.Sender).
		WillReturnError(fmt.Errorf("there is an error"))

	mock.ExpectRollback()
	err := repository.Transfer(transfer)
	if assert.NotNil(t, err) {
		assert.Equal(t, "there is an error", err.Error())
	}
}

func TestReceiveTransferFailure(t *testing.T) {
	db, mock := NewDBMock()
	repository := &Repositories{db}

	defer repository.DB.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(utils.SENDTRANSFER)).WithArgs(transfer.Amount, transfer.Sender).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta(utils.RECEIVETRANSFER)).WithArgs(transfer.Amount, transfer.Receiver).WillReturnError(fmt.Errorf("there is an error"))
	mock.ExpectRollback()

	err := repository.Transfer(transfer)

	if assert.NotNil(t, err) {
		assert.Equal(t, "there is an error", err.Error())
	}

}
