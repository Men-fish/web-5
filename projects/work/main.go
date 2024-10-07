package main

import (
	"fmt"
	"sync"
)

func work() {
	fmt.Println("Work is done")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			work()
		}()

		wg.Wait()

	}
	fmt.Println("All work is done")
}
