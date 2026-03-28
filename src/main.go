package main

import "fmt"

func inputData() float64 {
	var a, b float64
	var znak string

	fmt.Print("Input left operand: ")
	fmt.Scan(&a)
	fmt.Print("Input operation: ")
	fmt.Scan(&znak)
	fmt.Print("Input right operand:")
	fmt.Scan(&b)

	switch znak {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			fmt.Print("Invalid input\n")
			return inputData()
		} else {
			return a / b
		}
	default:
		fmt.Print("Invalid input\n")
		return inputData()
	}

}

func main() {
	result := inputData()
	fmt.Printf("%.3f\n", result)
}
