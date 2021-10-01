package main
import "fmt"

func main() {
	// Note: The value of a constant must be assigned when you declare it
	const name string = "kok dara"
	fmt.Println("Name: ", name)

	const sex = "Male"
	fmt.Println("Sex: ", sex)

	// error
	/*
		const position string = "Developer"
		position = "Cleaner" // error Constants: Unchangeable and Read-only
	*/

	// Multiple Constants Declaration
	const(
		A int = 1000
		B float32 = 3.14
		C string = "Hello"
		D bool = true
	)
	fmt.Println("A: ", A, "B: ", B, "C: ", C, "D: ", D)
}