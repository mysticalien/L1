package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает подаваемую на вход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

// Вариант 1: Метод с использованием двух указателей
func reverse1(s string) string {
	runes := []rune(s) // Преобразуем строку в массив рун для корректной обработки Unicode символов
	n := len(runes)

	for i := 0; i < n/2; i++ {
		// Меняем местами элементы: первый с последним, второй с предпоследним и так далее
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

// Вариант 2: Запись в обратном порядке
func reverse2(s string) string {
	runes := []rune(s)
	res := make([]rune, len(runes))
	for i, j := range runes {
		res[len(res)-1-i] = j
	}
	return string(res)
}

// Вариант 3: Используя strings.Builder
func reverse3(s string) string {
	var builder strings.Builder
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i]) // Добавляет один символ (руну) к текущему содержимому builder
	}
	return builder.String() // Преобразует текущее содержимое builder в строку и возвращает его
}

// Вариант 4: Рекурсивный метод
// Может привести к переполнению стека для очень длинных строк из-за большого числа рекурсивных вызовов
func reverse4(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	return string(runes[len(runes)-1]) + reverse4(string(runes[:len(runes)-1]))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Original:", input)
	fmt.Println("Reversed (Method 1):", reverse1(input))
	fmt.Println("Reversed (Method 2):", reverse2(input))
	fmt.Println("Reversed (Method 3):", reverse3(input))
	fmt.Println("Reversed (Method 4):", reverse4(input))
}
