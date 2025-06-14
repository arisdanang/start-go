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

	fmt.Println("-------------")

	var bAnimals = animals[0:2]
	// append new element
	var cAnimals = append(bAnimals, "fish")
	fmt.Println(animals)
	fmt.Println(cAnimals)
	fmt.Println(bAnimals)

	fmt.Println("-------------")

	dst := make([]string, 3)
	src := []string{"bird","dog","fish","lion"}
	n := copy(dst,src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	fmt.Println("-------------")


	// access slice using 3 indexes
	var people = []string{"Budi","Agung","Irwan"}
	var aPeople = people[0:2]
	var bPeople = people[0:2:2]

	fmt.Println(people)
	fmt.Println(len(people))
	fmt.Println(cap(people))

	fmt.Println(aPeople)
	fmt.Println(len(aPeople))
	fmt.Println(cap(aPeople))


	fmt.Println(bPeople)
	fmt.Println(len(bPeople))
	fmt.Println(cap(bPeople))



}
