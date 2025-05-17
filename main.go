package main

import (
	"fmt"
)

func main(){
	// var nama string = "danang"

	// var status string
	// status = "single"

	// type inference
	// job := "software engineer"
	// job = "frontend engineer"

	// declare multi var
	// var pertama, kedua, ketiga string
	// pertama, kedua, ketiga = "first", "second", "third"

	// declare multi var with type inference
	// keempat, kelima, keenam := "fourth", "fifth", "sixth"

	// declare multi var with type inference and different type
	// one, isHoliday, pi, greeting := 1, true, 3.14, "hi"


	// declare var with new keyword
	// name := new(string)

	// fmt.Println(*name)


	// fmt.Println("Hello world", "icik bos")
	// fmt.Printf("hello %s %s %s \n", nama, status, job)
	// fmt.Printf("%s %s %s \n", pertama, kedua, ketiga)
	// fmt.Printf("%s %s %s \n", keempat, kelima, keenam)

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

	// konstanta
	const firstName string  = "Aris"

	fmt.Print("halo ", firstName, "! \n")

	fmt.Println("aris danang")
	fmt.Println("aris","danang")

	fmt.Print("aris danang\n")
	fmt.Print("aris ","danang\n")
	fmt.Print("aris", " ", "danang\n")

	// deklarasi multi konstanta
	const (
		square = "persegi"
		isExist bool = true
		numeric uint8 = 23
		floatNum = 2.3
	)

	fmt.Println(square,isExist,numeric, floatNum)

	const (
		a = "testing"
		b
	)

	fmt.Println(a,b)

	const satu, dua = 1, 2
	const three, four, five string = "tiga", "empat", "lima"

	fmt.Println(satu,dua,three,four,five)

	// conditional statement
	var trafficLight = "green"

	if trafficLight == "red" {
		fmt.Println("berhenti")
	} else if trafficLight == "yellow" {
		fmt.Println("bersiap siap")
	} else {
		fmt.Println("jalan!!!")
	}


	// temp variable inside conditional statements
	var point = 1777.3

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect! \n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good! \n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad!\n", percent, "%")
	}

	// switch case
	var score = 9

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7,6,5,4:
		fmt.Println("awesome")
	default: {
		fmt.Println("not bad")
		fmt.Println("you can be better")
	}
	}

	// switch case dengan gaya if-else
	var age uint8 = 17;

	switch {
	case age >= 17 :
		fmt.Println("dewasa")
		fallthrough // keyword yg unik dan membingungkan hehehehe
	case age < 17:
		fmt.Println("anak kecil")
	case age < 5:
		fmt.Println("balita")
	default:
		fmt.Println("batita")
	}


}
