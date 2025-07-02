package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func main() {
	var data interface{}

	data = 2
	var result = data.(int) * 10
	fmt.Println(data, "multiplied by 10 is :", result)

	data = []string{"running", "music", "swimming"}
	var hobbies = strings.Join(data.([]string), ", ")
	fmt.Println(hobbies, " is my favorite hobbies")

	var mahasiswa interface{} = &person{name: "danang", age: 22}
	name := mahasiswa.(*person).name
	fmt.Println(name)

}
