package main

import (
	"fmt"
	"time"
)

func main() {
	go printName("arjun")
	go printName("arya")
	time.Sleep(time.Millisecond * 500)
}

func printName(name string) {
	fmt.Println(name)
}
