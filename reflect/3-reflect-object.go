package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	reflectValue := reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	reflectType := reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama : ", reflectType.Field(i).Name)
		fmt.Println("tipe data : ", reflectType.Field(i).Type)
		fmt.Println("nilai : ", reflectValue.Field(i).Interface())
		fmt.Println("=====")
	}
}

func main() {
	student1 := &student{Name: "danang", Grade: 23}
	student1.getPropertyInfo()
}
