package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int

		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}

		return min, max
	}

	numbers := []int{1, 2, 3, 4, 5, 6}
	min, max := getMinMax(numbers)
	fmt.Printf("data: %v\nmin: %v\nmax: %v\n", numbers, min, max)

	var newNumbers = func(min int) []int {
		var r []int

		for _, e := range numbers {
			fmt.Println(e)
			if e < min {
				continue
			}

			r = append(r, e)
		}

		return r
	}(3)

	fmt.Println("original number:", numbers)
	fmt.Println("filtered number:", newNumbers)

	fmt.Println("------------")

	howMany, getNumbers := findMax([]int{3, 2, 4, 6, 5, 3, 2, 1, 5}, 4)
	// nums := getNumbers()
	fmt.Println("data: \t", howMany)
	fmt.Println("value: \t", getNumbers())

}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}
