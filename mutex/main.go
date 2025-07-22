package main

import (
	"fmt"
	"sync"
)

// In this example mutex help to ge concorrency.
func main() {
	fmt.Println("ener the number")
	var wg sync.WaitGroup
	var lk sync.Mutex
	counter := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, &lk, &counter)
	}
	wg.Wait()
	fmt.Println(counter)
}

func increment(wg *sync.WaitGroup, lk *sync.Mutex, counter *int) {
	defer wg.Done()
	lk.Lock()
	*counter++
	lk.Unlock()
}
