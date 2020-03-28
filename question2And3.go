package main

import "fmt"

func main() {
	fmt.Print("Enter First String: ")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Print("Enter Second String: ")
	var num2 int
	fmt.Scanln(&num2)

	defer Add(num1, num2)
	defer Subtract(num1, num2)
	defer Multiply(num1, num2)
	defer Divide(num1, num2)
	fmt.Println("Calculator:-")
}
func Add(num1, num2 int) {
	println("Addition: ", num1+num2)
}

func Subtract(num1, num2 int) {
	println("Subtract: ", num1-num2)
}

func Multiply(num1, num2 int) {
	println("Multiplication: ", num1*num2)
}

func Divide(num1, num2 int) {
	println("Division: ", num1+num2)
}
