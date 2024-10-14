package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func swapArithmetic(a, b int) (int, int) {
	a += b    // 38
	b = a - b // 17
	a -= b    // 21
	return a, b
}

// Используя исключающее И (Результат 1, если биты различаются)
func swapUsingXOR(a, b int) (int, int) {
	a = a ^ b // 10001 ^ 10101 = 00100
	b = a ^ b // 00100 ^ 10101 = 10001
	a = a ^ b // 00100 ^ 10001 = 10101
	return a, b
}

func main() {
	a, b := 17, 21

	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)

	a, b = swap(a, b)
	fmt.Printf("After swap (Method 1): a = %d, b = %d\n", a, b)

	a, b = swapArithmetic(a, b)
	fmt.Printf("After swap (Method 2): a = %d, b = %d\n", a, b)

	a, b = swapUsingXOR(a, b)
	fmt.Printf("After swap (Method 3): a = %d, b = %d\n", a, b)
}
