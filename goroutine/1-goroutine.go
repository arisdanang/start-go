package main

import (
	"fmt"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	// runtime.GOMAXPROCS(2)

	// go print(5, "halo")
	// print(5, "how are you?")

	// var input string
	// fmt.Scanln(&input)

	var w1, w2, w3 string
	fmt.Scanln(&w1, &w2, &w3)

	fmt.Println(w1)
	fmt.Println(w2)
	fmt.Println(w3)

}
