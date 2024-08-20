package pointer_error

import (
	"errors"
	"fmt"
)

type Bitcoin int
type ErrorFound error

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance Bitcoin
}

type WalletErr struct {
	err ErrorFound
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) WalletErr {
	if amount > w.balance {
		return WalletErr{
			err: ErrInsufficientFunds,
		}
	}

	w.balance -= amount
	return WalletErr{}
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (we *WalletErr) Error() string {
	return fmt.Sprintf("%s", we.err.Error())
}
