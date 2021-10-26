package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}

func ExampleSum() {
	sum := Sum(5, 5)
	fmt.Println(sum)
	// Output: 10
}
