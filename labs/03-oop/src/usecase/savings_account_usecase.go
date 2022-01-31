package usecase

import (
	"github.com/fabianogoes/learning-golang/labs/03-oop/src/domain/account"
	"github.com/fabianogoes/learning-golang/labs/03-oop/src/infra/integration"
)

func SavingsAccountDoDeposit(sa *account.SavingsAccount, amount float64) {
	_, err := sa.Deposit(amount)
	if err != nil {
		panic(err)
	}
}

func SavingsAccountDoWithdraw(sa *account.SavingsAccount, amount float64) {
	_, err := sa.Withdraw(amount)
	if err != nil {
		panic(err)
	}
}

func SavingsAccountDoTransfer(source *account.SavingsAccount, amount float64, target *account.SavingsAccount) {
	_, err := source.Transfer(amount, target)
	if err != nil {
		panic(err)
	}
}

func SavingsAccountGetBalance(sa *account.SavingsAccount) float64 {
	return sa.GetBalance()
}

func SavingsAccountPay(sa *account.SavingsAccount, amountDocument float64) {
	_, err := sa.Withdraw(amountDocument)
	if err != nil {
		panic(err)
	}

	integration.DoPayment(amountDocument)
}
