package main

import "fmt"

func main() {
	var data interface{}

	data = "hello world"
	fmt.Println(data)

	data = 12
	fmt.Println(data)

	data = []string{"fish", "bird", "lion"}
	fmt.Println(data)

}
