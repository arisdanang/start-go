package main

import "fmt"

func main() {
	defer fmt.Println("hallo")
	fmt.Println("selamat datang")

	fmt.Println("=======")

	orderFood("pizza")
	orderFood("rawon")

}

func orderFood(menu string) {
	defer fmt.Println("terima kasih, silahkan tunggu")

	if menu == "pizza" {
		fmt.Print("pilihan tepat", " ")
		fmt.Print("pizza ditempat kami sangat enak", "\n")
		return
	}

	fmt.Println("pesanan kamu: ", menu)
}
