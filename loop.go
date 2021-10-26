package main

import "strings"

func Loop(str string, cycles int) (result string) {
	for i := 0; i < cycles; i++ {
		result += str
	}
	return
}

func CompareStr(str1, str2 string) int {
	return strings.Compare(str1, str2)
}
