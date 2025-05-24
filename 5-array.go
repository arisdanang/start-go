package main

import "fmt"


func main(){
	var players [4]string
	players[0] = "ronaldo"
	players[1] = "messi"
	players[2] = "rooney"
	players[3] = "haaland"

	fmt.Println(players[0])
	fmt.Println("jumlah element: ",len(players))
	fmt.Println("isi semua element: ", players)

	fmt.Println("-------------")

	// cara 2 - horizontal
	var animals [4]string

	animals = [4]string{"bird","dog","cat","fish"}

	fmt.Println(animals)

	fmt.Println("-------------")


	// cara 3 - vertikal

	animals = [4]string {
		"tiger",
		"shark",
		"lion",
		"horse",
	}

	fmt.Println("isi semua element: ", animals)
	fmt.Println("jumlah element: ", len(animals))

	fmt.Println("-------------")

	// cara 3 - init array tanpa size
	var countries = [...]string{"arab","indonesia", "singapore","malaysia"}

	fmt.Println("isi semua element \t: ", countries)
	fmt.Println("jumlah semua element \t: ", len(countries) )

}
