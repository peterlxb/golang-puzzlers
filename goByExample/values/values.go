package main

import "fmt"
/*
 Go has various value types
*/
func main() {
	// string
	fmt.Println("go"+"lang")

	// integer
	fmt.Println("1+1 =", 1 + 1)

	// floats
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// boolean
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}