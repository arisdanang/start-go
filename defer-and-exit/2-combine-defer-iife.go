package main

import "fmt"

func main() {
	number := 3

	if number == 3 {
		fmt.Println("halo 1")
		// jika kita mau defer dieksekusi di akhir blok if, kita bisa bungkus menggunakan IIFE
		func() {
			defer fmt.Println("halo 3")
		}()
	}

	fmt.Println("halo 2")
}
