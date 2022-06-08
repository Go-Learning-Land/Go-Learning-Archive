package main

import "fmt"

func main() {

	m := make(map[string]int) // To create an empty map, use the builtin make: make(map[key-type]val-type).

	m["k1"] = 7 //Set key/value pairs using typical name[key] = val syntax.
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"] //The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2} //	Note that maps appear in the form map[k:v k:v] when printed with fmt.Println.
	fmt.Println("map:", n)
}
