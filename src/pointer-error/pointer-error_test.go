package pointer_error

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got WalletErr, want string) {
		t.Helper()

		if got.err == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got.Error(), want)
		}
	}

	t.Run("wallet get balance", func(t *testing.T) {
		wallet := Wallet{}
		want := Bitcoin(0)
		assertBalance(t, wallet, want)
	})

	t.Run("wallet deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(100))
		want := Bitcoin(100)
		assertBalance(t, wallet, want)
	})

	t.Run("wallet Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(100))
		err := wallet.Withdraw(Bitcoin(85))
		want := Bitcoin(15)

		if err.err != nil {
			assertError(t, err, "cannot withdraw, insufficient funds")
		}

		assertBalance(t, wallet, want)
	})

	t.Run("wallet Withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{startBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, startBalance)
	})
}
