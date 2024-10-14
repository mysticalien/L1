package main

import "fmt"

/*
Дана переменная int64.
Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func setBit(n int64, i int, bitValue int) int64 {
	if bitValue == 1 {
		return n | (1 << i) // '|' — побитовое OR
	} else {
		return n &^ (1 << i) // '&^' — побитовое AND NOT
	}
}

func main() {
	fmt.Print("Enter a number (int64): ")
	var num int64
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid int64 number.")
		return
	}

	fmt.Print("Enter the bit index (0-63): ")
	var i int
	_, err = fmt.Scan(&i)
	if err != nil || i < 0 || i > 63 {
		fmt.Println("Invalid bit index. Please enter a value between 0 and 63.")
		return
	}

	fmt.Print("Enter the bit value (1 or 0): ")
	var bitValue int
	_, err = fmt.Scan(&bitValue)
	if err != nil || (bitValue != 1 && bitValue != 0) {
		fmt.Println("Invalid bit value. Please enter 1 or 0.")
		return
	}

	num = setBit(num, i, bitValue)
	fmt.Printf("Result after setting the %d-th bit to %d: %d\n", i, bitValue, num)
}
