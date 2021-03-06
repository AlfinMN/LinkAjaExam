package utils

const (
	CHECKSALDO      = `SELECT acc.account_number,cus.customer_name as customer_name,acc.balance from account acc join customer cus on acc.customer_number = cus.customer_number WHERE acc.account_number=?`
	SENDTRANSFER    = `UPDATE account set balance = balance - ?  where account_number = ?`
	RECEIVETRANSFER = `UPDATE account set balance = balance + ? where account_number = ?`
)
