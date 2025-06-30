package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	person1 := struct {
		person
		grade int
	}{}

	person1.person = person{"danang", 22}
	person1.grade = 1

	fmt.Println("name : ", person1.person.name)
	fmt.Println("name : ", person1.person.age)
	fmt.Println("name : ", person1.grade)

}
