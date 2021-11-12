package main

import "fmt"

func main() {
	numbersChan := make(chan int)
	go func() {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for _, number := range numbers {
			numbersChan <- number
		}
		close(numbersChan) //will stop for range loop
	}()

	for number := range numbersChan { //stops on close()
		fmt.Println(number)
	}
}
