package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSumArray(t *testing.T) {
	t.Run("test a variable length array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		result := SumArray(numbers)

		expected := 10
		if expected != result {
			t.Errorf("esperado %d, obtido %d, data %v", expected, result, numbers)
		}
	})
}

func BenchmarkSumArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumArray([]int{b.N})
	}
}

func ExampleSumArray() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(SumArray(numbers))
	//OutPut 15
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2, 3}, []int{4, 5, 6})

	expected := []int{6, 15}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("esperado %v, obtido %v", expected, result)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumArray(make([]int, b.N))
	}
}

func ExampleSumAll() {
	SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	fmt.Println([]int{6, 15})
	//OutPut {6, 15}
}

func TestSumEverythingElse(t *testing.T) {
	verifyResult := func(t *testing.T, expected, result []int) {
		t.Helper()
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("esperado %v, obtido %v", expected, result)
		}
	}

	t.Run("test sum everything else of array", func(t *testing.T) {
		result := SumEverythingElse([]int{1, 2, 3}, []int{4, 5, 6})
		expected := []int{5, 11}
		verifyResult(t, expected, result)
	})

	t.Run("test a sum for empty array", func(t *testing.T) {
		result := SumEverythingElse([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		verifyResult(t, expected, result)
	})
}

func BenchmarkSumEverythingElse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumEverythingElse(make([]int, b.N))
	}
}

func ExampleSumEverythingElse() {
	SumEverythingElse([]int{}, []int{3, 4, 5})
	fmt.Println([]int{0, 9})
	//OutPut {0, 9}
}
