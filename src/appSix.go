package main
import "fmt"

func main() {

	/*
		%v is used to print the value of the arguments
		%T is used to print the type of the arguments
	*/
	var number int = 100
	fmt.Printf("Number has value: %v and type: %T", number, number)
}