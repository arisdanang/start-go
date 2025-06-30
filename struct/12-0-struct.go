package main

import "fmt"

type mahasiswa struct {
	nama string
	nim  string
}

func main() {
	var mahasiswa1 mahasiswa
	mahasiswa1.nama = "danang"
	mahasiswa1.nim = "22222"

	var mahasiswa2 = mahasiswa{}
	mahasiswa2.nama = "adam"
	mahasiswa2.nim = "123123"

	var mahasiswa3 = mahasiswa{"dewi", "23223423"}

	var mahasiswa4 = mahasiswa{nama: "dini", nim: "2323423"}

	var mahasiswa5 = mahasiswa{nim: "345345", nama: "ari"}

	fmt.Println("mahasiswa 1, name \t:", mahasiswa1.nama)

	fmt.Println("mahasiswa 1, nim \t:", mahasiswa1.nim)

	fmt.Println("mahasiswa 2, name \t:", mahasiswa2.nama)
	fmt.Println("mahasiswa 2, nim \t:", mahasiswa2.nim)

	fmt.Println("mahasiswa 3, name \t:", mahasiswa3.nama)
	fmt.Println("mahasiswa 3, nim \t:", mahasiswa3.nim)

	fmt.Println("mahasiswa 4, name \t:", mahasiswa4.nama)
	fmt.Println("mahasiswa 4, nim \t:", mahasiswa4.nim)

	fmt.Println("mahasiswa 5, name \t:", mahasiswa5.nama)
	fmt.Println("mahasiswa 5, nim \t:", mahasiswa5.nim)

	// embedded struct

}
