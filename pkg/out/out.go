package out

import "fmt"

func Greet() {
	fmt.Println("Please choose the type of die you want to roll roll:")
	fmt.Println("Enter 'a' to rol a D4")
	fmt.Println("Enter 'b' to rol a D6")
	fmt.Println("Enter 'c' to rol a D8")
	fmt.Println("Enter 'd' to rol a D20")
	fmt.Println("Enter 'e' to rol a D100")
	fmt.Println("Enter 'q' to quit")
}
