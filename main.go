package main

import (
	"fmt"
	"simple-calculator/calculator"
)

func main() {

	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	var result float64
	var err error

	switch operator {
	case "+":
		result = calculator.Add(num1, num2)
	case "-":
		result = calculator.Substract(num1, num2)
	case "*":
		result = calculator.Multiply(num1, num2)
	case "/":
		result, err = calculator.Divide(num1, num2)
		if err != nil {
			fmt.Println("❌", err)
			return
		}
	default:
		fmt.Println("❌ Invalid operator.")
		return
	}

	fmt.Printf("✅ Result: %.2f\n", result)
}


