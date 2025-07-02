package main

import "fmt"

func main() {
	var data map[string]any

	data = map[string]any{
		"name":    "john",
		"grade":   12,
		"hobbies": []string{"running", "music", "swimming"},
	}

	fmt.Println(data)
}
