package internal

import "testing"

func TestService(t *testing.T) {
	service := NewAccountBusiness()

	t.Run("find customer by cpf", func(t *testing.T) {
		if c1 := service.FindCustomerByCpf(customerCpf123); c1 == nil {
			t.Errorf("should be found a customer with cpf: %s", customerCpf123)
		}
		if c1 := service.FindCustomerByCpf(customerCpf456); c1 == nil {
			t.Errorf("should be found a customer with cpf: %s", customerCpf456)
		}
	})

	t.Run("deposit", func(t *testing.T) {
		// Given

		depositAmount := 10.5
		expectedBalance := service.RetrieveBalance(accountNumber1) + depositAmount

		// When
		transactionId := service.Deposit(accountNumber1, depositAmount)

		// Then
		if updatedBalance := service.FindAccountByNumber(accountNumber1).RetrieveBalance(); updatedBalance != expectedBalance {
			t.Errorf("should be balance updated to: %v after deposit but balance is: %v", expectedBalance, updatedBalance)
		}
		if !service.ExistTransaction(transactionId, TransactionTypeDeposit) {
			t.Errorf("should be exist a transation id: %v", transactionId)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		// Given
		depositAmount := 100.
		service.Deposit(accountNumber1, depositAmount)

		// When
		withdrawAmount := 10.5
		expectedBalance := service.RetrieveBalance(accountNumber1) - withdrawAmount
		transactionId := service.Withdraw(accountNumber1, withdrawAmount)

		// Then
		if updatedBalance := service.FindAccountByNumber(accountNumber1).RetrieveBalance(); updatedBalance != expectedBalance {
			t.Errorf("should be balance updated to: %v after withdraw but balance is: %v", expectedBalance, updatedBalance)
		}
		if !service.ExistTransaction(transactionId, TransactionTypeWithdraw) {
			t.Errorf("should be exist a transation id: %v", transactionId)
		}
	})

	t.Run("payment", func(t *testing.T) {
		// Given
		depositAmount := 100.
		service.Deposit(accountNumber1, depositAmount)

		// When
		paymentAmount := 10.5
		expectedBalance := service.RetrieveBalance(accountNumber1) - paymentAmount
		transactionId := service.Pay(accountNumber1, "123", paymentAmount)

		// Then
		if updatedBalance := service.FindAccountByNumber(accountNumber1).RetrieveBalance(); updatedBalance != expectedBalance {
			t.Errorf("should be balance updated to: %v after pay but balance is: %v", expectedBalance, updatedBalance)
		}
		if !service.ExistTransaction(transactionId, TransactionTypePayment) {
			t.Errorf("should be exist a transation id: %v", transactionId)
		}
	})

	t.Run("transfer", func(t *testing.T) {
		// Given
		depositAmount := 100.
		service.Deposit(accountNumber1, depositAmount)

		// When
		transferAmount := 10.5
		expectedBalanceTarget := service.RetrieveBalance(accountNumber2) + transferAmount
		expectedBalanceSource := service.RetrieveBalance(accountNumber1) - transferAmount
		transactionId := service.Transfer(accountNumber1, accountNumber2, transferAmount)

		// Then
		if updatedBalance := service.FindAccountByNumber(accountNumber2).RetrieveBalance(); updatedBalance != expectedBalanceTarget {
			t.Errorf("should be balance updated to: %v after transfer but balance is: %v", expectedBalanceTarget, updatedBalance)
		}
		if updatedBalance := service.FindAccountByNumber(accountNumber1).RetrieveBalance(); updatedBalance != expectedBalanceSource {
			t.Errorf("should be balance updated to: %v after transfer but balance is: %v", expectedBalanceSource, updatedBalance)
		}
		if !service.ExistTransaction(transactionId, TransactionTypeTransfer) {
			t.Errorf("should be exist a transation id: %v", transactionId)
		}
	})
}
