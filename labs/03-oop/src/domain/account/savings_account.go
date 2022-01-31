package account

import (
	"fmt"

	"github.com/fabianogoes/learning-golang/labs/03-oop/src/domain/customer"
)

type SavingsAccount struct {
	ID       int
	Customer customer.Customer
	Agency   string
	Account  string
	balance  float64
}

func (sa *SavingsAccount) Withdraw(amount float64) (newBalance float64, err error) {
	fmt.Printf("Account: %v, Withdraw: %v\n", sa.Account, amount)
	if amount < 0 {
		return sa.balance, fmt.Errorf("invalid amount: %v", amount)
	}

	if amount > sa.balance {
		return sa.balance, fmt.Errorf("insufficient balance: %v to withdraw: %v", sa.balance, amount)
	}

	sa.balance -= amount
	return sa.balance, nil
}

func (sa *SavingsAccount) Deposit(amount float64) (newBalance float64, err error) {
	fmt.Printf("Account: %v, Deposit: %v\n", sa.Account, amount)
	if amount < 0 {
		return sa.balance, fmt.Errorf("invalid amount: %v", amount)
	}

	sa.balance += amount
	return sa.balance, nil
}

func (a *SavingsAccount) GetBalance() float64 {
	return a.balance
}

func (sa *SavingsAccount) Transfer(amount float64, target *SavingsAccount) (newBalance float64, err error) {
	fmt.Printf("Account: %v, Transfer: %v, Target: %v\n", sa.Account, amount, target.Account)
	if amount < 0 {
		return sa.balance, fmt.Errorf("invalid amount to transfer: %v", amount)
	}

	if sa.balance < amount {
		return sa.balance, fmt.Errorf("insufficient balance: %v to transfer: %v", sa.balance, amount)
	}

	sa.Withdraw(amount)
	target.Deposit(amount)
	return sa.balance, nil
}
