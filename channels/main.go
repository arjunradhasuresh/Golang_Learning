package main

import (
	"fmt"
	"time"
)

func main() {
	var num1 int
	fmt.Println("enter the number")
	fmt.Scan(&num1)
	fmt.Println("number is", num1)
	chan1 := make(chan int, 3)
	chan1 <- num1

	go add(chan1)

	time.Sleep(time.Millisecond * 500)
}

func add(cha1 chan int) {
	v1 := <-cha1
	fmt.Println(v1 + 5)
}
