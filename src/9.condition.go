package main

import (
	"fmt"
)

func main() {
	// example of using if statement
	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	x := 20
	y := 18
	if x > y {
		fmt.Println("x is greater than y")
	}

	// example of using if else statement
	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	// time := 20
	// if time < 18 {
	// 	fmt.Println("Good day.")
	// } // this raises an error
	// else {
	// 	fmt.Println("Good evening.")
	// }

	// example of using if else if statement
	a := 14
	b := 14
	if a < b {
		fmt.Println("a is less than b.")
	} else if a > b {
		fmt.Println("a is more than b.")
	} else {
		fmt.Println("a and b are equal.")
	}

	// example of using nested if statement
	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}
}
