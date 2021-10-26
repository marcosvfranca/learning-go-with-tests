package main

import "testing"

func TestHello(t *testing.T) {
	checkMessage := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	t.Run("show hello with name", func(t *testing.T) {
		resultado := Hello("Marcos", "")
		esperado := "Hello, Marcos"
		checkMessage(t, resultado, esperado)
	})
	t.Run("return hello, world where empty string", func(t *testing.T) {
		resultado := Hello("", "")
		esperado := "Hello, World"
		checkMessage(t, resultado, esperado)
	})
	t.Run("with lang espanhol", func(t *testing.T) {
		resultado := Hello("Marcos", "espanhol")
		esperado := "Ei, Marcos"
		checkMessage(t, resultado, esperado)
	})
	t.Run("with lang frances", func(t *testing.T) {
		resultado := Hello("Marcos", "frances")
		esperado := "Hi, Marcos"
		checkMessage(t, resultado, esperado)
	})
}
