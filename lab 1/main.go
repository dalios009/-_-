package main

import (
	"errors"
	"fmt"
)

// Step 2: Function `hello` that takes a string and returns a greeting message.
func hello(name string) string {
	return "Привет, " + name + "!"
}

// Step 3: Function `printEven` that prints even numbers in a given range.
func printEven(a, b int64) error {
	if a > b {
		return errors.New("левая граница больше правой границы диапазона")
	}
	fmt.Printf("Четные числа в диапазоне от %d до %d: ", a, b)
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println() // Newline for better formatting
	return nil
}

// Step 4: Function `apply` that performs arithmetic operations based on an operator.
func apply(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль невозможно")
		}
		return a / b, nil
	default:
		return 0, errors.New("действие не поддерживается")
	}
}

func main() {
	// Step 2: Test the `hello` function
	fmt.Println(hello("Мир"))

	// Step 3: Test the `printEven` function
	err := printEven(2, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	err = printEven(10, 2) // This should return an error
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	// Step 4: Test the `apply` function
	result, err := apply(3, 5, "+")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("3 + 5 = %.2f\n", result)
	}

	result, err = apply(7, 10, "*")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("7 * 10 = %.2f\n", result)
	}

	result, err = apply(3, 5, "#") // Unsupported operation
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Result = %.2f\n", result)
	}

	result, err = apply(3, 0, "/") // Division by zero
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("3 / 0 = %.2f\n", result)
	}
}
