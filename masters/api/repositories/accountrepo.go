package repositories

import "TestLinkAja/masters/api/models"

type AccountRepo interface {
	CheckSaldo(string) (*models.Account, error)
	Transfer(*models.Transfer) error
}
