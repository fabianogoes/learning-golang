package usecase

import (
	"github.com/fabianogoes/learning-golang/labs/03-oop/src/domain/account"
	"github.com/fabianogoes/learning-golang/labs/03-oop/src/infra/integration"
)

func CheckingAccountDoDeposit(ca *account.CheckingAccount, amount float64) {
	_, err := ca.Deposit(amount)
	if err != nil {
		panic(err)
	}
}

func CheckingAccountDoWithdraw(ac *account.CheckingAccount, amount float64) {
	_, err := ac.Withdraw(amount)
	if err != nil {
		panic(err)
	}
}

func CheckingAccountDoTransfer(caSource *account.CheckingAccount, amount float64, caTarget *account.CheckingAccount) {
	_, err := caSource.Transfer(amount, caTarget)
	if err != nil {
		panic(err)
	}
}

func CheckingAccountGetBalance(ca *account.CheckingAccount) float64 {
	return ca.GetBalance()
}

func CheckingAccountDoPay(ca *account.CheckingAccount, amountDocument float64) {
	_, err := ca.Withdraw(amountDocument)
	if err != nil {
		panic(err)
	}

	integration.DoPayment(amountDocument)
}
