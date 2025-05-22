package main

import "fmt"

func main(){
	var name string = "danang"

	fmt.Println(name)

	// type inference
	job := "software engineer"

	fmt.Println(job)

	// deklarasi multi variable
	var pertama, kedua, ketiga = "first", "kedua", "ketiga"

	// deklarasi multi var with type inference
	empat, lima, enam := "four", "five", "six"

	fmt.Println(pertama,kedua,ketiga,empat,lima,enam)


	// deklarasi multi var dengan type inference dengan berbeda type data
	one, isHoliday, pi, greeting := 1, true, 3.14, "hi"

	fmt.Println(one,isHoliday,pi, greeting)

}
