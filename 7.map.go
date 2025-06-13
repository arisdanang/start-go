package main

import "fmt"

func main(){

	// horizontal style
	var days = map[string]int{"senin": 10, "selasa": 20, "rabu": 30}

	// vertikal
	var days2 = map[string]int{
		"senin": 12,
		"selasa": 14,
		"rabu": 16,
		"kamis": 18,
	}


	fmt.Println("senin", days["senin"])
	fmt.Println("selasa", days2["selasa"])

	// *init map without default value
	// var days3 = map[string]int{}
	// var day4 = make(map[string]int)
	// var days5 = *new(map[string]int)


	// looping with for...range
	for key, val := range days2 {
		fmt.Println(key, " \t:", val)
	}

	// delete item
	delete(days2, "kamis")
	fmt.Println(days2)

	// find existence item
	var value, isExist = days2["rabu"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item not exists")
	}

	// combine slice and map
	var cats = []map[string]string {
		map[string]string{"name": "larry","gender": "male"},
		map[string]string{"name": "candy","gender": "female"},
	}

	for _, cat := range cats {
		fmt.Println(cat["name"], cat["gender"])
	}

	// *other ways to combine slice and map
	// var cats = []map[string]string {
	// 	{"name": "larry","gender": "male"},
	// 	{"name": "candy","gender": "female"},
	// }

	// we can add item with different keys
	var data = []map[string]string{
		{"name": "John", "age": "21"},
		{"job": "software engineer", "status": "single"},
	}
}
