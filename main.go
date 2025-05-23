package main

import (
	"fmt"
)

func main(){
	// declare var with new keyword
	// name := new(string)

	// fmt.Println(*name)


	// fmt.Println("Hello world", "icik bos")
	// fmt.Printf("hello %s %s %s \n", nama, status, job)
	// fmt.Printf("%s %s %s \n", pertama, kedua, ketiga)
	// fmt.Printf("%s %s %s \n", keempat, kelima, keenam)


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
