package repositories

import (
	"TestLinkAja/masters/api/models"
	"TestLinkAja/masters/utils"
	"database/sql"
	"log"
)

type Repositories struct {
	DB *sql.DB
}

func InitAcountRepository(db *sql.DB) AccountRepo {
	return &Repositories{DB: db}
}

// Check Saldo Repo

func (r *Repositories) CheckSaldo(accountNumber string) (*models.Account, error) {

	account := &models.Account{}

	row := r.DB.QueryRow(utils.CHECKSALDO, accountNumber)

	err := row.Scan(&account.AccountNumber, &account.CustomerName, &account.Balance)

	if err != nil {
		return nil, err
	}

	return account, nil

}

// Transfer Repo

func (r *Repositories) Transfer(transfer *models.Transfer) error {

	tx, err := r.DB.Begin()

	if err != nil {
		log.Println(err)
		return err
	}

	_, err = tx.Exec(utils.SENDTRANSFER, &transfer.Amount, &transfer.Sender)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	_, err = tx.Exec(utils.RECEIVETRANSFER, &transfer.Amount, &transfer.Receiver)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	tx.Commit()

	return nil
}
