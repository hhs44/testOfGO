package main

import (
	"errors"
	"fmt"
	"testing"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}
type Stringer interface {
	String() string
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
func (w *Wallet) Deposit(i Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += i
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
func (w *Wallet) WithDraw(i Bitcoin) error {
	if i > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= i
	return nil
}



func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t *testing.T, got  error, want string) {
		t.Helper()
		if got == nil {
        	t.Fatal("didn't get an error but wanted one")
		}
		if got.Error() != want {
			t.Errorf("got %q,want %q", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("WithDraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(100)}
		wallet.WithDraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(90))
	})
	t.Run("WithDraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.WithDraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, "cannot withdraw, insufficient funds")
	})

}
