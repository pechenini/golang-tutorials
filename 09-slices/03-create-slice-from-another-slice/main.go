package main

import "fmt"

func main() {
	cities := []string{"New York", "London", "Chicago", "Beijing", "Delhi", "Hong Kong"}

	asianCities := cities[3:]

	fmt.Println("cities = ", cities)
	fmt.Println("asianCities = ", asianCities)
}
