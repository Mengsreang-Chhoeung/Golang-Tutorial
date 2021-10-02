package main

import (
	"fmt"
)

func main() {
	// example: access elements of a slice
	prices := []int{10, 20, 30}
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// example: append elements to a slice
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// example: append one slice to another slice
	myslice2 := []int{1, 2, 3}
	myslice3 := []int{4, 5, 6}
	myslice4 := append(myslice2, myslice3...)

	fmt.Printf("myslice4=%v\n", myslice4)
	fmt.Printf("length=%d\n", len(myslice4))
	fmt.Printf("capacity=%d\n", cap(myslice4))
}
