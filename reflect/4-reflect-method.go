package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	student1 := &student{Name: "Aris Danang", Grade: 22}
	fmt.Println("Nama :", student1.Name)

	reflectValue := reflect.ValueOf(student1)

	method := reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Gugun"),
	})

	fmt.Println("Nama :", student1.Name)

}
