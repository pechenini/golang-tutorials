package main

import "fmt"

func main() {
	// The append() function appends new elements at the end of a given slice.

	// Appending to a slice that doesn't have enough capacity to accommodate new elements
	slice1 := []string{"C", "C++", "Java"}
	slice2 := append(slice1, "Python", "Ruby", "Go")

	fmt.Printf("slice1 = %v, len = %d, cap = %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 = %v, len = %d, cap = %d\n", slice2, len(slice2), cap(slice2))

	slice1[0] = "C#"
	fmt.Println("\nslice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
}
