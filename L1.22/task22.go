package main

import (
	"fmt"
	"math/big"
	"os"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
*/

func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func subtract(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func divide(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("Error: Division by zero")
		os.Exit(1)
	}
	return new(big.Int).Div(a, b)
}

func main() {
	a := new(big.Int)
	b := new(big.Int)

	var inputA, inputB string
	var op string

	fmt.Print("Enter first number (a): ")
	_, err := fmt.Scan(&inputA)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Print("Enter second number (b): ")
	_, err = fmt.Scan(&inputB)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Установка значений в big.Int
	a.SetString(inputA, 10)
	b.SetString(inputB, 10)

	// Можно добавть проверку на то, введено ли число больше 2^20

	//if a.Cmp(big.NewInt(1048576)) <= 0 || b.Cmp(big.NewInt(1048576)) <= 0 {
	//	fmt.Println("Both numbers must be greater than 2^20.")
	//	os.Exit(1)
	//}

	fmt.Print("Enter operation (+, -, *, /): ")
	_, err = fmt.Scan(&op)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	var result *big.Int
	switch op {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result = divide(a, b)
	default:
		fmt.Println("Invalid operation")
		os.Exit(1)
	}

	fmt.Printf("Result: %s\n", result.String())
}
