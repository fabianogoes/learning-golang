package internal

import (
	"github.com/google/uuid"
)

const (
	customerCpf123 = "123"
	customerCpf456 = "456"

	accountNumber1 = 1
	accountNumber2 = 2
)

var customerList = []Customer{
	{Id: uuid.New(), Name: "Customer 1", Cpf: customerCpf123},
	{Id: uuid.New(), Name: "Customer 2", Cpf: customerCpf456},
}

var accountList = []*Account{
	{Id: uuid.New(), Number: accountNumber1, Customer: &customerList[0]},
	{Id: uuid.New(), Number: accountNumber2, Customer: &customerList[1]},
}
var transactionList []Transaction

type memoryDatabase struct {
}

func NewMemoryDB() *memoryDatabase {
	return &memoryDatabase{}
}

func (mdb *memoryDatabase) FindCustomerByCpf(cpf string) *Customer {
	for _, item := range customerList {
		if item.Cpf == cpf {
			return &item
		}
	}
	return nil
}

func (mdb *memoryDatabase) FindAccountByNumber(number int) *Account {
	for _, item := range accountList {
		if item.Number == number {
			return item
		}
	}
	return nil
}

func (mdb *memoryDatabase) UpdateAccount(account *Account) *Account {
	for index, item := range accountList {
		if item.Number == account.Number {
			accountList[index] = account
			return mdb.FindAccountByNumber(account.Number)
		}
	}
	return nil
}

func (mdb *memoryDatabase) InsertTransaction(transaction Transaction) uuid.UUID {
	transactionList = append(transactionList, transaction)
	return transaction.Id
}

func (mdb *memoryDatabase) ListAllTransactions() []Transaction {
	return transactionList
}

func (mdb *memoryDatabase) ExistTransaction(id uuid.UUID, withType string) bool {
	for _, t := range transactionList {
		if t.Id == id {
			return t.transactionType == withType
		}
	}
	return false
}
