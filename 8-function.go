package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main()  {
	var names = []string{"danang", "wahyu"}
	greetHello("hallo", names)

	var randomValue int

	randomValue = randomWithRange(2,10)
	fmt.Println("random", randomValue)

	randomValue = randomWithRange(2,10)
	fmt.Println("random", randomValue)

	randomValue = randomWithRange(2,10)
	fmt.Println("random", randomValue)

}

func greetHello(message string, arr []string){
	var name = strings.Join(arr, " ")
	fmt.Println(message, name)
}

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}
