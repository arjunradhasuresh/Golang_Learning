package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//here we are using wait group which is used to use only specific amount of users can only run rest need to wait()
	//sync package is using
	//wg.add()
	//wg.Done()
	//wg.sleep() are using

	var wg sync.WaitGroup
	var num int

	fmt.Println("enter the number of goroutine you need to create")
	fmt.Scan(&num)

	for i := 0; i <= num; i++ {
		wg.Add(1)
		go add(&wg, i)
		time.Sleep(time.Millisecond * 500)
	}

	wg.Wait()

	fmt.Println("compleated")
}
func add(wg *sync.WaitGroup, num int) {
	fmt.Println(num + 5)
	defer wg.Done()
}
