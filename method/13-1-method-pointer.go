package main

import "fmt"

type mahasiswa struct {
	name string
	nim  int
}

func (m mahasiswa) changeName1(name string) {
	fmt.Println("----> on changeName1, name changed to", name)

	m.name = name
}

func (m *mahasiswa) changeName2(name string) {
	fmt.Println("----> on changeName2, name changed to", name)
	m.name = name
}

func main() {
	mahasiswa1 := mahasiswa{"Rafi ahmad", 23}
	fmt.Println("mahasiswa1 before", mahasiswa1.name)

	mahasiswa1.changeName1("abdul hadi")
	fmt.Println("mahasiswa1 after changeName1", mahasiswa1.name)

	mahasiswa1.changeName2("haris julio")
	fmt.Println("mahasiswa1 after changeName2", mahasiswa1.name)
}
