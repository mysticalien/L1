package main

import (
	"fmt"
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}
func main() {
  someFunc()
}
*/

// Негативные последствия:
// 1. Утечка памяти: созданный срез ссылается на весь оригинальный массив `v`,
// поэтому неиспользуемая часть строки остается в памяти.
// 2. Проблемы со сборщиком мусора: поскольку срез ссылается на оригинальный массив,
// сборщик мусора не сможет освободить память, занимаемую неиспользуемой частью,
// что может привести к увеличению использования памяти.

// Решение:
// Создание копии строки вместо сохранения среза. Это позволяет сборщику мусора
// освободить память, занимаемую неиспользуемой частью, поскольку новая строка
// не ссылается на оригинальный массив.

var justString string

func createHugeString(size int) string {
	return strings.Repeat("o", size)
}

func someFuncWithString() {
	v := createHugeString(1 << 10) // Создаем строку размером 1024 байта
	justString = string(v[:100])   // Создаем копию первых 100 символов
}

func someFuncWithCopy() {
	v := createHugeString(1 << 10)
	buf := make([]byte, 100) // Создаем новый срез
	copy(buf, v[:100])       // Копируем первые 100 байт
	justString = string(buf) // Преобразуем срез байтов обратно в строку
}

func main() {
	someFuncWithString()
	someFuncWithCopy()
	fmt.Println(justString)
}
