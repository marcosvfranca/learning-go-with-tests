package main

func SumArray(numbers []int) (soma int) {
	for _, number := range numbers {
		soma += number
	}

	return
}

func SumAll(numbers ...[]int) (sum []int) {
	for _, number := range numbers {
		sum = append(sum, SumArray(number))
	}
	return
}

func SumEverythingElse(numbers ...[]int) (sum []int) {
	for _, number := range numbers {
		if len(number) == 0 {
			sum = append(sum, 0)
		} else {
			final := number[1:]
			sum = append(sum, SumArray(final))
		}
	}
	return
}
