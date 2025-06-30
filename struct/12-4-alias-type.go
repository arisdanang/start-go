package main

import "fmt"

type Person struct {
	name string
	age  int
}
type People = Person

func main() {

	p1 := Person{"wick", 21}
	fmt.Println(p1)

	p2 := Person{"niki", 22}
	fmt.Println(p2)

	people := People{"surya", 22}
	fmt.Println(Person(people))

	person := Person{"surya", 22}
	fmt.Println(People(person))

}
