package main

import "fmt"

//. From this point on we will be writing tests first. go back test file
// Hello returns a personalised greeting.
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
