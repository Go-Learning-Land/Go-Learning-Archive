package main

import (
	"fmt"
	"math"
)

const s string = "constant" //declare a constant

const ( // declaration of multiple constants
	licenseKey    = "9508-7466-2492-4708-9399"
	versionNumber = "1.0.0"
	licenseType   = "MIT"
)

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	fmt.Printf("%v - %v - %v \n", licenseKey, versionNumber, licenseType) // For details look for https://gobyexample.com/string-formatting

}
