package main

import "fmt"

func main() {
	var avg = calculate(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("rata-rata: %.2f", avg)
	var numbers = []int{1, 2, 3, 4}
	var total = add(numbers...)
	fmt.Printf("total : %d \n", total)
	fmt.Println(add(1, 2))
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var average = float64(total) / float64(len(numbers))

	return average
}

func add(numbers ...int) int {
	var total int = 0

	for _, number := range numbers {
		total += number
	}

	return total
}
