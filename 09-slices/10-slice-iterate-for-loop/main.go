package main

import "fmt"

func main() {
	countries := []string{"Ukraine", "Germany", "Italy", "England"}

	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}

	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	//for range
	for index, number := range primeNumbers {
		fmt.Printf("PrimeNumber(%d) = %d\n", index+1, number)
	}
}
