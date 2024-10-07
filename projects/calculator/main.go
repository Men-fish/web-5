package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resultChan := make(chan int)
	go func() {
		defer close(resultChan)
		select {
		case val := <-firstChan:
			resultChan <- val * val
		case val := <-secondChan:
			resultChan <- val * 3
		case <-stopChan:
			return
		}
	}()
	return resultChan
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	result := calculator(firstChan, secondChan, stopChan)
	go func() {
		//firstChan <- 7
		secondChan <- 8
		close(firstChan)
		close(secondChan)
		close(stopChan)
	}()
	for rest := range result {
		fmt.Println(rest)
	}

}
