package app

import (
	"fmt"

	"github.com/fabianogoes/learning-golang/labs/03-oop/src/infra/database"
	"github.com/fabianogoes/learning-golang/labs/03-oop/src/usecase"
)

func ExecuteSavingsAccount() {
	fmt.Println("-----------------------------------------")
	fmt.Println(">>> Savings Account Transactions")
	fmt.Println("-----------------------------------------")

	savingAccount1 := database.FindSavingsAccountById(3)
	savingAccount2 := database.FindSavingsAccountById(4)

	fmt.Printf("Account: %v, Balance: %v\n", savingAccount1.Account, usecase.SavingsAccountGetBalance(&savingAccount2))
	usecase.SavingsAccountDoDeposit(&savingAccount1, 1000)
	usecase.SavingsAccountDoWithdraw(&savingAccount1, 50)
	usecase.SavingsAccountDoTransfer(&savingAccount1, 100, &savingAccount2)
	usecase.SavingsAccountPay(&savingAccount1, 100)

	fmt.Println("-----------------------------------------")
	fmt.Printf("Account: %v, Balance: %v\n", savingAccount1.Account, usecase.SavingsAccountGetBalance(&savingAccount1))
	fmt.Printf("Account: %v, Balance: %v\n", savingAccount1.Account, usecase.SavingsAccountGetBalance(&savingAccount2))
}
