package models

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuiteTestAccount struct {
	suite.Suite
	Account
}
type SuiteTestTransfer struct {
	suite.Suite
	Transfer
}

func (suite *SuiteTestAccount) SetupTest() {
	suite.Account.AccountNumber = "5550010"
	suite.Account.CustomerName = "Odegaard"
	suite.Account.Balance = 3000
}

func (suite *SuiteTestTransfer) SetupTest() {
	suite.Transfer.Receiver = "5550010"
	suite.Transfer.Sender = "5550011"
	suite.Transfer.Amount = 200
}

func (suite *SuiteTestAccount) TestAccount() {
	suite.Equal(suite.Account.AccountNumber, "5550010")
	suite.Equal(suite.Account.CustomerName, "Odegaard")
	suite.Equal(suite.Account.Balance, 3000)
}

func (suite *SuiteTestTransfer) TestTransfer() {
	suite.Equal(suite.Transfer.Receiver, "5550010")
	suite.Equal(suite.Transfer.Sender, "5550011")
	suite.Equal(suite.Transfer.Amount, 200)
}

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(SuiteTestAccount))
	suite.Run(t, new(SuiteTestTransfer))
}
