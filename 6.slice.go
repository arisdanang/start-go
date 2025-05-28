package main

import "fmt"

func main(){
	// init slice
	var animals = []string{"dog","bird","cat","horse"}
	fmt.Println(animals[2])

	fmt.Println("-------------")

	var newAnimals = animals[0:3]
	fmt.Println(newAnimals)
	fmt.Println("-------------")

	// get all element
	fmt.Println(animals[:])
	fmt.Println("-------------")

	// get element from index 2 and so on
	fmt.Println(animals[2:])
	fmt.Println("-------------")

	// get element before index 2
	fmt.Println(animals[:2])
	fmt.Println("-------------")

	// get length of slice
	fmt.Println(len(animals))




}
