package pointAndErr

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		//fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		//fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		assertBalance(t, wallet, Bitcoin(10))
	})

}
