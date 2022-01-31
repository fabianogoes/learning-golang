package app

import (
	"fmt"

	"github.com/fabianogoes/learning-golang/labs/03-oop/src/infra/database"
	"github.com/fabianogoes/learning-golang/labs/03-oop/src/usecase"
)

func ExecuteCheckingAccount() {
	fmt.Println("-----------------------------------------")
	fmt.Println(">>> Checking Account Transactions")
	fmt.Println("-----------------------------------------")

	checkingAccount1 := database.FindCheckingAccountById(1)
	checkingAccount2 := database.FindCheckingAccountById(2)

	fmt.Printf("Account: %v, Balance: %v\n", checkingAccount1.Account, usecase.CheckingAccountGetBalance(&checkingAccount1))
	usecase.CheckingAccountDoDeposit(&checkingAccount1, 1000)
	usecase.CheckingAccountDoWithdraw(&checkingAccount1, 50)
	usecase.CheckingAccountDoTransfer(&checkingAccount1, 100, &checkingAccount2)
	usecase.CheckingAccountDoPay(&checkingAccount1, 100)

	fmt.Println("-----------------------------------------")
	fmt.Printf("Account: %v, Balance: %v\n", checkingAccount1.Account, usecase.CheckingAccountGetBalance(&checkingAccount1))
	fmt.Printf("Account: %v, Balance: %v\n", checkingAccount2.Account, usecase.CheckingAccountGetBalance(&checkingAccount2))
}
