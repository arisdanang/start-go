package main

import (
	"fmt"
	"strings"
)

type mahasiswa struct {
	name string
	nim  int
}

func (m mahasiswa) sayHello() {
	fmt.Println("halo ", m.name)
}

func (m mahasiswa) getNameAt(i int) string {
	return strings.Split(m.name, " ")[i-1]
}

func main() {
	mahasiswa1 := mahasiswa{"Aris Danang", 22}
	mahasiswa1.sayHello()

	nama := mahasiswa1.getNameAt(2)
	fmt.Println("nama panggilan : ", nama)
}
