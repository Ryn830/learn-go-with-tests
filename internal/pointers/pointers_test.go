package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		expected := Bitcoin(10)

		assertBalance(t, wallet, expected)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		withdrawErr := wallet.Withdraw(Bitcoin(10))
		expected := Bitcoin(10)

		assertNoError(t, withdrawErr)
		assertBalance(t, wallet, expected)
	})

	t.Run("Insufficient funds", func(t *testing.T) {
		startingFunds := Bitcoin(20)
		wallet := Wallet{startingFunds}
		insufficientFundsError := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingFunds)
		assertError(t, insufficientFundsError, InsufficientFundsError)
	})
}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()

	actual := wallet.Balance()
	if actual != expected {
		t.Errorf("%#v, Actual %s not equal %s", wallet, actual, expected)
	}
}

func assertError(t *testing.T, err error, expectedError error) {
	t.Helper()

	if err == nil {
		t.Fatal("Expected an error to be returned")
	}

	if err != expectedError {
		t.Errorf("Received %q, expected error message to be: %q", err, expectedError)
	}
}

func assertNoError(t *testing.T, returnedError error) {
	t.Helper()

	if returnedError != nil {
		t.Errorf("Received %q when no error is expected", returnedError)
	}
}
