package main

import (
	"fmt"
)

func main() {
	// example with using length
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	fmt.Println(arr1)
	fmt.Println(arr2)

	// example without using length
	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{4, 5, 6, 7, 8}
	fmt.Println(arr3)
	fmt.Println(arr4)

	// example: array of strings
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)

	// example: access elements of an array
	prices := [3]int{10, 20, 30}
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// example: change elements of an array
	prices[2] = 50
	fmt.Println(prices)
}
