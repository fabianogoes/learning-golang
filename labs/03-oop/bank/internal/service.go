package internal

import (
	"strconv"

	"github.com/google/uuid"
)

type AccountBusiness struct {
	customerRepo    CustomerRepository
	accountRepo     AccountRepository
	transactionRepo TransactionRepository
}

func NewAccountBusiness() *AccountBusiness {
	customerRepository := NewCustomerRepository()
	return &AccountBusiness{
		customerRepo:    customerRepository,
		accountRepo:     NewAccoutRepository(customerRepository),
		transactionRepo: NewTransactionRepository(),
	}
}

func (ab *AccountBusiness) FindCustomerByCpf(cpf string) *Customer {
	return ab.customerRepo.FindByCpf(cpf)
}

func (ab *AccountBusiness) FindAccountByNumber(number int) *Account {
	return ab.accountRepo.FindAccountByNumber(number)
}

func (ab *AccountBusiness) RetrieveBalance(accountNumber int) float64 {
	return ab.accountRepo.FindAccountByNumber(accountNumber).RetrieveBalance()
}

func (ab *AccountBusiness) Deposit(accountNumber int, amount float64) uuid.UUID {
	account := ab.accountRepo.FindAccountByNumber(accountNumber)
	if account == nil {
		panic("Account:" + strconv.Itoa(accountNumber) + " not found")
	}

	account.Deposit(amount)

	ab.accountRepo.UpdateAccount(account)
	return ab.transactionRepo.InsertTransaction(Transaction{Id: uuid.New(), Account: account, Amount: amount, transactionType: TransactionTypeDeposit})
}

func (ab *AccountBusiness) Withdraw(accountNumber int, amount float64) uuid.UUID {
	account := ab.accountRepo.FindAccountByNumber(accountNumber)
	if account == nil {
		panic("Account:" + strconv.Itoa(accountNumber) + " not found")
	}

	account.Withdraw(amount)
	ab.accountRepo.UpdateAccount(account)
	return ab.transactionRepo.InsertTransaction(Transaction{Id: uuid.New(), Account: account, Amount: amount, transactionType: TransactionTypeWithdraw})
}

func (ab *AccountBusiness) Pay(accountNumber int, barCode string, amount float64) uuid.UUID {
	account := ab.accountRepo.FindAccountByNumber(accountNumber)
	if account == nil {
		panic("Account:" + strconv.Itoa(accountNumber) + " not found")
	}

	// Send to payment

	account.Withdraw(amount)
	ab.accountRepo.UpdateAccount(account)
	return ab.transactionRepo.InsertTransaction(Transaction{Id: uuid.New(), Account: account, Amount: amount, transactionType: TransactionTypePayment})
}

func (ab AccountBusiness) Transfer(sourceAccountNumber int, targetAccountNumber int, amount float64) uuid.UUID {
	source := ab.accountRepo.FindAccountByNumber(sourceAccountNumber)
	if source == nil {
		panic("Source Account:" + strconv.Itoa(sourceAccountNumber) + " not found")
	}

	ab.Withdraw(sourceAccountNumber, amount)
	ab.Deposit(targetAccountNumber, amount)
	return ab.transactionRepo.InsertTransaction(Transaction{Id: uuid.New(), Account: nil, Amount: amount, transactionType: TransactionTypeTransfer})
}

func (ab AccountBusiness) ExistTransaction(id uuid.UUID, withType string) bool {
	return ab.transactionRepo.ExistTransaction(id, withType)
}
