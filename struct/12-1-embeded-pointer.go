package main

import "fmt"

type person struct {
	name string
	age  int
}

type mahasiswa struct {
	grade int
	age   int
	person
}

func main() {
	mahasiswa1 := mahasiswa{}
	mahasiswa1.name = "johan"
	mahasiswa1.age = 11
	mahasiswa1.grade = 1
	mahasiswa1.person.age = 44

	mahasiswa1.person.name = "ari"

	fmt.Println("name :", mahasiswa1.name)
	fmt.Println("grade :", mahasiswa1.grade)
	fmt.Println("age :", mahasiswa1.person.age)
	fmt.Println("age :", mahasiswa1.age)
	fmt.Println("name :", mahasiswa1.name)

}
