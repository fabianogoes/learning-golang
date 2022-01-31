package account

import (
	"fmt"

	"github.com/fabianogoes/learning-golang/labs/03-oop/src/domain/customer"
)

type CheckingAccount struct {
	ID       int
	Customer customer.Customer
	Agency   string
	Account  string
	balance  float64
}

func (ca *CheckingAccount) Withdraw(amount float64) (newBalance float64, err error) {
	fmt.Printf("Account: %v, Withdraw: %v\n", ca.Account, amount)
	if amount < 0 {
		return ca.balance, fmt.Errorf("invalid amount to withdraw: %v", amount)
	}

	if amount > ca.balance {
		return ca.balance, fmt.Errorf("insufficient balance: %v to withdraw: %v", ca.balance, amount)
	}

	ca.balance -= amount
	return ca.balance, nil
}

func (ca *CheckingAccount) Deposit(amount float64) (newBalance float64, err error) {
	fmt.Printf("Account: %v, Deposit: %v\n", ca.Account, amount)
	if amount < 0 {
		return ca.balance, fmt.Errorf("invalid amount to deposit: %v", amount)
	}

	ca.balance += amount
	return ca.balance, nil
}

func (a *CheckingAccount) GetBalance() float64 {
	return a.balance
}

func (ca *CheckingAccount) Transfer(amount float64, target *CheckingAccount) (newBalance float64, err error) {
	fmt.Printf("Account: %v, Transfer: %v, Target: %v\n", ca.Account, amount, target.Account)
	if amount < 0 {
		return ca.balance, fmt.Errorf("invalid amount to transfer: %v", amount)
	}

	if ca.balance < amount {
		return ca.balance, fmt.Errorf("insufficient balance: %v to transfer: %v", ca.balance, amount)
	}

	ca.Withdraw(amount)
	target.Deposit(amount)
	return ca.balance, nil
}
