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

// var customerList = []Customer{
// 	{Id: uuid.New(), Name: "Customer 1", Cpf: customerCpf123},
// 	{Id: uuid.New(), Name: "Customer 2", Cpf: customerCpf456},
// }

// var accountList = []*Account{
// 	{Id: uuid.New(), Number: accountNumber1, Customer: &customerList[0]},
// 	{Id: uuid.New(), Number: accountNumber2, Customer: &customerList[1]},
// }
// var transactionList []Transaction

type CustomerRepository interface {
	FindByCpf(cpf string) *Customer
}

type customerDataAccess struct {
	list []*Customer
}

func (da *customerDataAccess) FindByCpf(cpf string) *Customer {
	for _, item := range da.list {
		if item.Cpf == cpf {
			return item
		}
	}
	return nil
}

func NewCustomerRepository() CustomerRepository {
	memoryList := []*Customer{
		{Id: uuid.New(), Name: "Customer 1", Cpf: customerCpf123},
		{Id: uuid.New(), Name: "Customer 2", Cpf: customerCpf456},
	}
	return &customerDataAccess{memoryList}
}

type AccountRepository interface {
	FindAccountByNumber(number int) *Account
	UpdateAccount(account *Account) *Account
}

type accountDataAccess struct {
	list []*Account
}

func NewAccoutRepository(customerRepo CustomerRepository) AccountRepository {
	memoryList := []*Account{
		{Id: uuid.New(), Number: accountNumber1, Customer: customerRepo.FindByCpf(customerCpf123)},
		{Id: uuid.New(), Number: accountNumber2, Customer: customerRepo.FindByCpf(customerCpf456)},
	}
	return &accountDataAccess{memoryList}
}

type memoryDatabase struct {
}

func NewMemoryDB() *memoryDatabase {
	return &memoryDatabase{}
}

func (ada *accountDataAccess) FindAccountByNumber(number int) *Account {
	for _, item := range ada.list {
		if item.Number == number {
			return item
		}
	}
	return nil
}

func (ada *accountDataAccess) UpdateAccount(account *Account) *Account {
	for index, item := range ada.list {
		if item.Number == account.Number {
			ada.list[index] = account
			return ada.FindAccountByNumber(account.Number)
		}
	}
	return nil
}

type TransactionRepository interface {
	InsertTransaction(transaction Transaction) uuid.UUID
	ListAllTransactions() []Transaction
	ExistTransaction(id uuid.UUID, withType string) bool
}

type transactionDataAccess struct {
	list []Transaction
}

func NewTransactionRepository() TransactionRepository {
	return &transactionDataAccess{}
}

func (tda *transactionDataAccess) InsertTransaction(transaction Transaction) uuid.UUID {
	tda.list = append(tda.list, transaction)
	return transaction.Id
}

func (tda *transactionDataAccess) ListAllTransactions() []Transaction {
	return tda.list
}

func (tda *transactionDataAccess) ExistTransaction(id uuid.UUID, withType string) bool {
	for _, t := range tda.list {
		if t.Id == id {
			return t.transactionType == withType
		}
	}
	return false
}
