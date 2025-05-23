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

	// konstanta
	const firstName string  = "Aris"

	fmt.Print("halo ", firstName, "! \n")

	// println => mencetak dengan spasi antar arg dan baris baru di akhir
	fmt.Println("aris danang")
	fmt.Println("aris","danang")

	// print => mencetak tanpa baris baru
	fmt.Print("aris danang\n")
	fmt.Print("aris","danang\n")
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



}
