package main

import "fmt"

func main() {
	ch := make(chan int)
	//pointered version //https://golangbyexample.com/channel-function-argument-go/

	go func(a, b int, ch *chan int) {
		c := a + b
		*ch <- c
	}(1, 2, &ch)
	c := <-ch
	fmt.Printf("computed value %v\n", c)
	// TODO: get the value computed from goroutine
	// fmt.Printf("computed value %v\n", c)
}
