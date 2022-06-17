//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {

	var slice = []int{1, 3, 5, 6}

	slice = append(slice, 8) //https://programming.guide/go/append-explained.html#:~:text=The%20built%2Din%20append%20function,the%20data%20is%20copied%20over.
	/*
		The built-in append function appends elements to the end of a slice:

		1) if there is enough capacity, the underlying array is reused;
		2) if not, a new underlying array is allocated and the data is copied over.
	*/

	fmt.Println("Capacity", cap(slice))

	fmt.Println("Length", len(slice))
}
