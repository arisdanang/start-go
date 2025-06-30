package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	allStudents := []person{
		{name: "jalal", age: 12},
		{name: "nabila", age: 22},
		{name: "danang", age: 12},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "umur : ", student.age)
	}
}
