package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false — если нет). Функция проверки должна быть
регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

// Функция проверки уникальности с использованием карты
func isUnique(s string) bool {
	s = strings.ToLower(s)

	// Если длина строки больше 52, уникальные символы невозможны (26 англ. + 26 рус.)
	if len(s) > 52 {
		return false
	}

	seen := make(map[rune]bool)

	for _, r := range s {
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

// Функция проверки уникальности с использованием битовой карты
func isUniqueBitMap(s string) bool {
	s = strings.ToLower(s)

	if len(s) > 52 {
		return false
	}

	var checker uint64

	for _, char := range s {
		var bitPosition uint64
		if char >= 'a' && char <= 'z' {
			bitPosition = uint64(char - 'a')
		} else if char >= 'а' && char <= 'я' {
			bitPosition = uint64(char - 'а' + 26) // Смещение на 26 для русского алфавита
		} else {
			continue
		}

		// Проверяем, установлен ли бит
		if checker&(1<<bitPosition) > 0 {
			return false
		}
		// Устанавливаем бит для данного символа
		checker |= 1 << bitPosition
	}
	return true
}

func main() {
	var str string
	fmt.Println("Enter a string: ")
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Using map")
	if isUnique(str) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	fmt.Println("Using bit map")
	if isUniqueBitMap(str) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
