package main
import "fmt"

func main() {
	var number int = 100 //type is int
	var floatNumber float32 = 123.25 //type is float32
	var text string = "Hello World" //type is string
	var isMarried bool = false //type is bool

	fmt.Println("1------------------------")
	fmt.Println("Number: ", number)
	fmt.Println("Floating Number: ", floatNumber)
	fmt.Println("Text: ", text)
	fmt.Println("Married: ", isMarried)
	fmt.Println("-------------------------")

	fmt.Print("\n")

	var numberTwo = 100 //type is inferred
	var floatNumberTwo = 123.25 //type is inferred
	var textTwo = "Hello World" //type is inferred
	var isMarriedTwo = false //type is inferred

	fmt.Println("2------------------------")
	fmt.Println("Number: ", numberTwo)
	fmt.Println("Floating Number: ", floatNumberTwo)
	fmt.Println("Text: ", textTwo)
	fmt.Println("Married: ", isMarriedTwo)
	fmt.Println("-------------------------")

	fmt.Print("\n")

	numberThree := 100 //type is inferred
	floatNumberThree := 123.25 //type is inferred
	textThree := "Hello World" //type is inferred
	isMarriedThree := false //type is inferred

	fmt.Println("3------------------------")
	fmt.Println("Number: ", numberThree)
	fmt.Println("Floating Number: ", floatNumberThree)
	fmt.Println("Text: ", textThree)
	fmt.Println("Married: ", isMarriedThree)
	fmt.Println("-------------------------")

	fmt.Print("\n")

	// default value when you initial variale without value
	var numberFour int 
	var floatNumberFour float32
	var textFour string 
	var isMarriedFour bool 

	fmt.Println("4------------------------")
	fmt.Println("Number: ", numberFour)
	fmt.Println("Floating Number: ", floatNumberFour)
	fmt.Println("Text: ", textFour)
	fmt.Println("Married: ", isMarriedFour)
	fmt.Println("-------------------------")

	// error when you initial varaible without value when using :=
	// number :=
	// and one more := can't declare outside main() function

	fmt.Print("\n")

	// multi-variable declaration
	var a, b, c, d int = 10, 20, 30, 40
	fmt.Println("A: ", a, "B: ", b, "C: ", c, "D: ", d)

	// Note: If you use the type keyword, it is only possible to declare one type of variable per line.

	fmt.Print("\n")

	// If the type of variable is not specified, you can declare different types of variables in the same line like below
	var aa, bb = 10, "Hello"
	var cc, dd = true, 19.99
	fmt.Println("AA: ", aa, "BB: ", bb, "CC: ", cc, "DD: ", dd)

	fmt.Print("\n")

	// Multiple variable declarations can also be grouped together into a block for greater readability
	var(
		aaa int = 100
		bbb string = "Hello"
		ccc float32 = 14.55
		ddd bool = true
	)
	fmt.Println("AAA: ", aaa, "BBB: ", bbb, "CCC: ", ccc, "DDD: ", ddd)

	/*
		Rules for Go variables:

		*	A variable name must start with a letter or the underscore character
		*	A variable name cannot start with a number
		*	A variable name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
		*	Variable names are case-sensitive (age, Age and AGE are three different variables)
		*	There is no limit on the length of the variable name
		*	A variable name cannot contain spaces
		*	The variable name cannot be any Go keywords
		*	To keep the code maintainable, do not use the name of built-in functions as variable names
	*/

	// Camel Case
	// myVariableName = "John"

	// Pascal Case
	// MyVariableName = "John"

	// Snake Case
	// my_variable_name = "John"
}