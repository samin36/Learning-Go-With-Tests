package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error, expectedError string) {
		t.Helper()

		if err == nil {
			t.Fatal("expected an error but didn't get one")
		}

		if err.Error() != expectedError {
			t.Errorf("got %q, but wanted %q", err, expectedError)
		}
	}

	assertNoError := func(t testing.TB, err error) {
		t.Helper()

		if err != nil {
			t.Errorf("expected no errors but received %q", err)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)

		assertBalance(t, wallet, startingBalance)
	})
}
