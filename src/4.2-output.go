package main
import "fmt"

func main() {
	var id int
	var name string
	var position string
	var salary float32

	fmt.Print("Input ID: ")
	fmt.Scanln(&id)
	// fmt.Scanf("%v", &id)
	fmt.Print("Input Name: ")
	// fmt.Scanln(&name)
	fmt.Scanf("%s", &name)
	fmt.Print("Input Position: ")
	// fmt.Scanln(&position)
	fmt.Scanf("%s", &position)
	fmt.Print("Input Salary: ")
	fmt.Scanln(&salary)
	// fmt.Scanf("%v", &salary)

	fmt.Println("ID: ", id, " Name: ", name, " Position: ", position, " Salary: ", salary , "$")
}