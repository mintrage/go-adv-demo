package main

import (
	"fmt"
	"time"
)

func main() {
	go printHi()
	go fmt.Println("Привет из main")
	go fmt.Println("Привет из main 2")
	time.Sleep(time.Second)
}

func printHi() {
	fmt.Println("Привет из gr")
}
