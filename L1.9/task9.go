package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func sendNumbers(numbers []int, outputChan chan<- int) {
	for _, num := range numbers {
		outputChan <- num
	}
	close(outputChan)
}

func multiplyNumbers(inputChan <-chan int, outputChan chan<- int) {
	for num := range inputChan {
		outputChan <- num * 2
	}
	close(outputChan)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	inputChan := make(chan int)
	outputChan := make(chan int)

	go sendNumbers(numbers, inputChan)
	go multiplyNumbers(inputChan, outputChan)

	fmt.Println("Results after multiplying by 2:")
	for result := range outputChan {
		fmt.Println(result)
	}
}
