package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

// Вариант 1: С использованием WaitGroup
func squareWaitGroup(numbers []int) {
	fmt.Println("*************************************")
	fmt.Println("Using WaitGroup")
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1) // Увеличиваем счетчик горутин
		go func(n int) {
			defer wg.Done() // Уменьшаем счетчик, когда горутина завершает работу
			fmt.Printf("Квадрат числа равен: %d\n", n*n)
		}(num)
	}

	wg.Wait() // Ждем завершения всех горутин
}

// Вариант 2: С использованием каналов
func squareChannel(numbers []int) {
	fmt.Println("*************************************")
	fmt.Println("Using Channel")
	ch := make(chan int)

	for _, num := range numbers {
		go func(n int) {
			ch <- n * n
		}(num)
	}

	for i := 0; i < len(numbers); i++ {
		result := <-ch
		fmt.Printf("Квадрат числа равен: %d\n", result)
	}

	close(ch)
}

// Вариант 3: С использованием WaitGroup и каналов
func squareChannelWaitGroup(numbers []int) {
	fmt.Println("*************************************")
	fmt.Println("Using Channel + WaitGroup")
	var wg sync.WaitGroup
	ch := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch <- n * n
		}(num)
	}

	wg.Wait()
	close(ch)

	for result := range ch {
		fmt.Printf("Квадрат числа равен: %d\n", result)
	}
}

// Вариант 4: Использование конвейера для вычисления
func squarePipeline(numbers []int) {
	fmt.Println("*************************************")
	fmt.Println("Using Pipeline")

	numberChan := make(chan int)
	resultChan := make(chan int)

	go func() {
		for _, num := range numbers {
			numberChan <- num
		}
		close(numberChan)
	}()

	go func() {
		for num := range numberChan {
			resultChan <- num * num
		}
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Printf("Квадрат числа равен: %d\n", result)
	}
}
func main() {
	numbers := []int{2, 4, 6, 8, 10}

	squareWaitGroup(numbers)

	squareChannel(numbers)

	squareChannelWaitGroup(numbers)

	squarePipeline(numbers)
}
