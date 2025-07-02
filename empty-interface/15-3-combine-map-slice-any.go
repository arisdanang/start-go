package main

import "fmt"

func main() {
	// combine slice, map, dan interface{}/any
	var person = []map[string]interface{}{
		{"name": "danang", "age": 23},
		{"name": "haikal", "age": 12},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

	fmt.Println("======")

	var fruits = []any{
		map[string]any{"name": "apple", "price": 10000},
		[]string{"mango", "orange", "jackfruit"},
		"watermelon",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}
}
