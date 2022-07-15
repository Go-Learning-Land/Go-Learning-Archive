//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt package
import (
	"fmt"
)

// main method
func main() {

	var arr = [5]int{1, 2, 4, 5, 6}

	var i int
	// You can use ==> for i:= 0; i < len(arr); i++ // Therefore you dont need to use ==> var i int
	for i = 0; i < len(arr); i++ {

		fmt.Println("printing elements ", arr[i])

	}

	var value int
	for i, value = range arr {

		fmt.Println(" range ", value)
		// Go Compiler accully warn us about "you dont use variable i" but we declared it in line 15

	}

	for _, value = range arr {
		// The _ blank identifier is used if the index is ignored
		fmt.Println("blank range", value)

	}

}
