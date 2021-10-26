package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		verifyBalance(t, wallet, 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{bankBalance: 20}
		err := wallet.Withdraw(10)
		verifyBalance(t, wallet, 10)
		confirmThereIsNoError(t, err)
	})

	t.Run("withdraw without bank balance", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		wallet := Wallet{bankBalance: saldoInicial}
		err := wallet.Withdraw(30)

		verifyBalance(t, wallet, saldoInicial)
		confirmErr(t, err, ErrInsuficientBankBalance)
	})
}

func verifyBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	result := wallet.Balance()
	if result != expected {
		t.Errorf("resultado %s, esperado %s", result, expected)
	}
}

func confirmErr(t *testing.T, result error, expected error) {
	t.Helper()
	if result == nil {
		t.Fatal("esperava um erro, mas nenhum ocorreu")
	}

	if result != expected {
		t.Errorf("err result %s, err expected %s", result, expected)
	}
}

func confirmThereIsNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("erro inesperado recebido")
	}
}
