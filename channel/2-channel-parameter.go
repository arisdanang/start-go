package main

import (
	"fmt"
	"runtime"
)

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	for _, each := range []string{"danang", "aris", "jalal"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			fmt.Println(each)
			messages <- data
		}("jan")
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}

	var pesan = make(chan string)

	go func(who string) {
		var data = fmt.Sprintf("hello, %s", who)
		pesan <- data
	}("danang")

	var message = <-pesan
	fmt.Println(message)
}
