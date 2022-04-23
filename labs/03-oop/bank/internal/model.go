package internal

import (
	"github.com/google/uuid"
)

type Customer struct {
	Id   uuid.UUID
	Name string
	Cpf  string
}

type Account struct {
	Id       uuid.UUID
	Number   int
	Customer *Customer
	balance  float64
}

func (a *Account) Deposit(amount float64) {
	a.balance = a.balance + amount
}

func (a *Account) Withdraw(amount float64) {
	if a.balance < amount {
		panic("Insufficient balance to withdraw")
	}
	a.balance = a.balance - amount
}

func (a *Account) RetrieveBalance() float64 {
	return a.balance
}

type Transaction struct {
	Id              uuid.UUID
	Account         *Account
	Amount          float64
	transactionType string
}

const (
	TransactionTypeDeposit  = "DEPOSIT"
	TransactionTypeWithdraw = "WITHDRAW"
	TransactionTypePayment  = "PAYMENT"
	TransactionTypeTransfer = "TRANSFER"
)
