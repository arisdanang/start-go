package main

import "fmt"

func main() {
	// tipe data - numerik non desimal
	var positifNum uint8 = 70
	var negativeNum = -1234523

	fmt.Printf("bilangan positif: %d\n", positifNum)
	fmt.Printf("bilangan negatif: %d\n", negativeNum)

	// tipe data - numerik desimal
	var decimalNum = 3.14

	fmt.Printf("bilangan desimal: %f\n", decimalNum)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNum)

	// tipe data - boolean
	var isOpen bool = true

	fmt.Printf("is it open?? %t\n", isOpen)

	// tipe data - string
	var message string = "hello world"

	fmt.Printf("pesan: %s \n", message)

	// using backtick
	var introduce = `Nama saya "Aris Danang"
	salam kenal
	mari kita belajar golang`

	fmt.Println(introduce)
}
