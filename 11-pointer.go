package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) \t:", numberA)
	fmt.Println("numberA (address) \t:", &numberA)
	fmt.Println("numberB (value) \t:", *numberB)
	fmt.Println("numberB (address) \t:", &numberB)

	number := 1
	fmt.Println("before \t:", number)

	change(&number, 11)
	fmt.Println("after \t:", number)

}

// pointer sebagai parameter
func change(original *int, value int) {
	*original = value
}
