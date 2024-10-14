package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun» становится «sun dog snow».
*/

// Переворачивает порядок слов в строке.
func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// Переворачивает слова с использованием регулярных выражений.
func reverseWordsRegexp(s string) string {
	re := regexp.MustCompile(`\S+`)
	words := re.FindAllString(s, -1)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func reverseWordsSplit(s string) string {
	words := strings.Split(s, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println("Enter a string:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)
	fmt.Println("Original:", input)
	fmt.Println("Reversed (Method 1):", reverseWords(input))
	fmt.Println("Reversed (Method 2):", reverseWordsRegexp(input))
	fmt.Println("Reversed (Method 3):", reverseWordsSplit(input))
}
