package main

import "fmt"

func main(){
	var nama string = "danang"

	var status string
	status = "single"

	// type inference
	job := "software engineer"
	job = "frontend engineer"

	// declare multi var
	var pertama, kedua, ketiga string
	pertama, kedua, ketiga = "first", "second", "third"

	// declare multi var with type inference
	keempat, kelima, keenam := "fourth", "fifth", "sixth"

	// declare multi var with type inference and different type
	// one, isHoliday, pi, greeting := 1, true, 3.14, "hi"


	// declare var with new keyword
	name := new(string)

	fmt.Println(*name)

	fmt.Println("Hello world", "icik bos")
	fmt.Printf("hello %s %s %s \n", nama, status, job)
	fmt.Printf("%s %s %s \n", pertama, kedua, ketiga)
	fmt.Printf("%s %s %s \n", keempat, kelima, keenam)
}
