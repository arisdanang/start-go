package main

import (
	"fmt"
	"os"
)

func main() {
	// meskipun defer diletakan sebelum exit(), maka statement defer tidak akan dieksekusi
	defer fmt.Println("halo")
	os.Exit(2)
	fmt.Println("selamat datang")
}
