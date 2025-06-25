package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("angka", i)
	}

	fmt.Println("---------------")
	// for loop hanya dengan kondisi
	var a = 0

	for a < 5 {
		fmt.Println("angka", a)
		a++
	}

	fmt.Println("---------------")

	// for loop tanpa argumen

	var b = 0

	for {
		fmt.Println("angka", b)

		b++

		if b == 5 {
			break
		}
	}

	fmt.Println("---------------")

	// for - range
	for i := range 5 {
		fmt.Print(i)
	}

	fmt.Println("\n---------------")

	// for range untuk string
	var xs = "123"

	for i, v := range xs {
		fmt.Println("index= ", i, "value= ", v)
	}

	fmt.Println("---------------")

	// penggunaan keyword break dan continue

	for i := 1; i <= 10; i++ {
		fmt.Println("luar ", i)
		if i%2 == 1 {
			fmt.Println("asdf ", i)

			continue
		}

		fmt.Println(i)

		if i > 8 {
			break
		}

		fmt.Println("angka", i)

	}

	fmt.Println("---------------")

	// perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
}
