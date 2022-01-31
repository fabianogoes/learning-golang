package database

import (
	"fmt"

	"github.com/fabianogoes/learning-golang/labs/03-oop/src/domain/account"
	"github.com/fabianogoes/learning-golang/labs/03-oop/src/domain/customer"
)

var customerId int
var accountId int
var accountNumber int

func customerNextId() int {
	customerId += 1
	return customerId
}

func accountNextId() int {
	accountId += 1
	return accountId
}

func nextAccountNumber() string {
	accountNumber += 1
	return fmt.Sprintf("%06d", accountNumber)
}

var customer_database []customer.Customer = []customer.Customer{
	{ID: customerNextId(), Name: "Fabiano", CPF: "123", Occupation: "Software Engineer"},
	{ID: customerNextId(), Name: "Outro", CPF: "321", Occupation: "Data Engineer"},
}

var account_checking_database []account.CheckingAccount = []account.CheckingAccount{
	{ID: accountNextId(), Customer: customer_database[0], Agency: "0001", Account: nextAccountNumber()},
	{ID: accountNextId(), Customer: customer_database[1], Agency: "0001", Account: nextAccountNumber()},
}

var account_savings_database []account.SavingsAccount = []account.SavingsAccount{
	{ID: accountNextId(), Customer: customer_database[0], Agency: "1001", Account: nextAccountNumber()},
	{ID: accountNextId(), Customer: customer_database[1], Agency: "1001", Account: nextAccountNumber()},
}

func FindCustomerById(id int) customer.Customer {
	for _, c := range customer_database {
		if c.ID == id {
			return c
		}
	}

	fmt.Printf("customer not found with ID: %v", id)
	return customer.Customer{}
}

func FindCheckingAccountById(id int) account.CheckingAccount {
	for _, a := range account_checking_database {
		if a.ID == id {
			return a
		}
	}

	fmt.Printf("checking account not found with ID: %v", id)
	return account.CheckingAccount{}
}

func FindSavingsAccountById(id int) account.SavingsAccount {
	for _, a := range account_savings_database {
		if a.ID == id {
			return a
		}
	}

	fmt.Printf("savings account not found with ID: %v", id)
	return account.SavingsAccount{}
}
