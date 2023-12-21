package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit Bitcoin in the wallet", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)

	})

	t.Run("withdraw Bitcoin from the wallet with funds", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(5))
		want := Bitcoin(5)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw Bitcoing from the wallet insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Wanting an error but didn't get one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got an error but not wanting one")
	}
}
