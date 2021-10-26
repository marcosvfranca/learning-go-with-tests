package main

import (
	"errors"
	"fmt"
)

var ErrInsuficientBankBalance = errors.New("cannot withdraw: insufficient balance")

type Bitcoin int

type Wallet struct {
	bankBalance Bitcoin
}

func (w *Wallet) Deposit(quantity Bitcoin) {
	w.bankBalance += quantity
}

func (w *Wallet) Withdraw(quantity Bitcoin) error {
	if quantity > w.bankBalance {
		return ErrInsuficientBankBalance
	}

	w.bankBalance -= quantity
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.bankBalance
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
