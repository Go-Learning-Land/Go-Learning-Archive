package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	for i := 0; i < 10; i++ {
		go worker(done)
	}
	//go worker(done)

	prs, e := <-done
	fmt.Println("prs:", prs, "e:", e)
}
