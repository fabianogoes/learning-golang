package internal

import (
	"github.com/google/uuid"
	"strconv"
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

type AccountBusiness struct {
	db *memoryDatabase
}

func NewAccountBusiness(mdb *memoryDatabase) *AccountBusiness {
	return &AccountBusiness{db: mdb}
}

func (ab *AccountBusiness) FindCustomerByCpf(cpf string) *Customer {
	return ab.db.FindCustomerByCpf(cpf)
}

func (ab *AccountBusiness) FindAccountByNumber(number int) *Account {
	return ab.db.FindAccountByNumber(number)
}

func (ab *AccountBusiness) RetrieveBalance(accountNumber int) float64 {
	return ab.db.FindAccountByNumber(accountNumber).RetrieveBalance()
}

func (ab *AccountBusiness) Deposit(accountNumber int, amount float64) uuid.UUID {
	account := ab.db.FindAccountByNumber(accountNumber)
	if account == nil {
		panic("Account:" + strconv.Itoa(accountNumber) + " not found")
	}

	account.Deposit(amount)

	ab.db.UpdateAccount(account)
	return ab.db.InsertTransaction(Transaction{Id: uuid.New(), Account: account, Amount: amount, transactionType: TransactionTypeDeposit})
}

func (ab *AccountBusiness) Withdraw(accountNumber int, amount float64) uuid.UUID {
	account := ab.db.FindAccountByNumber(accountNumber)
	if account == nil {
		panic("Account:" + strconv.Itoa(accountNumber) + " not found")
	}

	account.Withdraw(amount)
	ab.db.UpdateAccount(account)
	return ab.db.InsertTransaction(Transaction{Id: uuid.New(), Account: account, Amount: amount, transactionType: TransactionTypeWithdraw})
}

func (ab *AccountBusiness) Pay(accountNumber int, barCode string, amount float64) uuid.UUID {
	account := ab.db.FindAccountByNumber(accountNumber)
	if account == nil {
		panic("Account:" + strconv.Itoa(accountNumber) + " not found")
	}

	// Send to payment

	account.Withdraw(amount)
	ab.db.UpdateAccount(account)
	return ab.db.InsertTransaction(Transaction{Id: uuid.New(), Account: account, Amount: amount, transactionType: TransactionTypePayment})
}

func (ab AccountBusiness) Transfer(sourceAccountNumber int, targetAccountNumber int, amount float64) uuid.UUID {
	source := ab.db.FindAccountByNumber(sourceAccountNumber)
	if source == nil {
		panic("Source Account:" + strconv.Itoa(sourceAccountNumber) + " not found")
	}

	ab.Withdraw(sourceAccountNumber, amount)
	ab.Deposit(targetAccountNumber, amount)
	return ab.db.InsertTransaction(Transaction{Id: uuid.New(), Account: nil, Amount: amount, transactionType: TransactionTypeTransfer})
}

func (ab AccountBusiness) ExistTransaction(id uuid.UUID, withType string) bool {
	return ab.db.ExistTransaction(id, withType)
}
