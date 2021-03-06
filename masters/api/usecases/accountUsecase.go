package usecases

import "TestLinkAja/masters/api/models"

type AccountUsecases interface {
	CheckSaldo(string) (*models.Account, error)
	Transfer(*models.Transfer) error
}
