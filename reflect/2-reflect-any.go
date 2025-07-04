package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 9000

	reflectValue := reflect.ValueOf(number)

	fmt.Println("tipe variable : ", reflectValue.Type())
	fmt.Println("nilai variabel : ", reflectValue.Interface())

	fmt.Println("result", reflectValue.Interface().(int)+1)

}
