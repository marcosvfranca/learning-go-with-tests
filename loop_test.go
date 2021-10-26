package main

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	result := Loop("a", 5)
	expected := "aaaaa"

	if result != expected {
		t.Errorf("esperado %s, obtido %s", expected, result)
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop("a", b.N)
	}
}

func ExampleLoop() {
	result := Loop("a", 3)
	fmt.Println(result)
	// Output: aaa
}

func TestCompare(t *testing.T) {
	result := CompareStr("a", "a")
	expected := 0

	if result != expected {
		t.Errorf("esperado %d, recebido %d", expected, result)
	}
}
