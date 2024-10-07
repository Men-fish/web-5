package main

import "fmt"

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	defer close(outputStream)
	var prev string
	for first := true; ; first = false {
		if val, ok := <-inputStream; !ok {
			return
		} else if first || val != prev {
			outputStream <- val
			prev = val
		}
	}
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream)

	go func() {
		for _, val := range []string{"a", "яяяяяяя", "b", "b", "a", "c", "c", "c"} {
			inputStream <- val
		}
		close(inputStream)
	}()

	for val := range outputStream {
		fmt.Println(val)
	}
}
