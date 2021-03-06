package usecases

import (
	"TestLinkAja/masters/api/models"
	"TestLinkAja/masters/api/repositories"
	"errors"
)

type AccountUsecaseImpl struct {
	accountRepo repositories.AccountRepo
}

func InitAccountUsecase(accountRepo repositories.AccountRepo) AccountUsecases {
	return &AccountUsecaseImpl{accountRepo: accountRepo}
}

func (a *AccountUsecaseImpl) CheckSaldo(accountNumber string) (*models.Account, error) {

	account, err := a.accountRepo.CheckSaldo(accountNumber)

	if err != nil {
		return nil, errors.New("Account Not Found")
	}

	return account, nil
}
func (a *AccountUsecaseImpl) Transfer(transfer *models.Transfer) error {

	account, err := a.accountRepo.CheckSaldo(transfer.Sender)

	if err != nil {
		return err
	}

	if transfer.Amount > account.Balance {
		return errors.New("saldo tidak cukup")
	}
	err = a.accountRepo.Transfer(transfer)
	if err != nil {
		return err
	}
	return nil
}
