package main

import (
	"fmt"
)

func main() {
	// example of array initialization
	arr1 := [5]int{}              //not initialized
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	// example of initialize only specific elements
	arr4 := [5]int{1: 10, 2: 40}
	fmt.Println(arr4)

	// example of using len()
	arr5 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr6 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(arr5))
	fmt.Println(len(arr6))
}
