package main

import "fmt"

func main() {
	// example 1
	s6 := make([]int, 0)
	fmt.Printf("The capacity of s6: %d\n", cap(s6))
	for i := 1; i <= 5; i++ {
		s6 = append(s6, i)
		fmt.Printf("s6(%d): len: %d, cap: %d\n", i, len(s6), cap(s6))
	}
	fmt.Println()

	// example 2
	s7 := make([]int,1024)
	fmt.Printf("The capacity of s7: %d\n", cap(s7))
	fmt.Println()

	// example 3
	s8 := make([]int, 10)
	fmt.Printf("The capacity of s8: %d\n", cap(s8))
}

