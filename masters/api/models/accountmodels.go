package models

type Account struct {
	AccountNumber string `json:"account_number"`
	CustomerName  string `json:"customer_name"`
	Balance       int    `json:"balance"`
}

type Transfer struct {
	Receiver string ` json:"to_account_number"`
	Sender   string `json :sender`
	Amount   int    `json:"amount"`
}
