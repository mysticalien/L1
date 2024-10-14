package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42…) с использованием конкурентных вычислений.
*/

func printHeader(title string) {
	fmt.Println("\n_____________________________________")
	fmt.Printf("Using %s\n", title)
}

func printSum(sum int) {
	fmt.Printf("Сумма квадратов: %d\n", sum)
}

// Вариант 1: WaitGroup + Общая переменная с мьютексом
func sumSquaresWaitGroup(numbers []int) {
	printHeader("WaitGroup + Shared Variable (Mutex)")

	var wg sync.WaitGroup
	var sum int
	var mu sync.Mutex

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			sum += n * n
			mu.Unlock()
		}(num)
	}

	wg.Wait()
	printSum(sum)
}

// Вариант 2: Канал для передачи результатов
func sumSquaresChannel(numbers []int) {
	printHeader("Channel")

	ch := make(chan int)

	for _, num := range numbers {
		go func(n int) {
			ch <- n * n
		}(num)
	}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += <-ch
	}

	close(ch)

	printSum(sum)
}

// Вариант 3: WaitGroup + Канал
func sumSquaresChannelWaitGroup(numbers []int) {
	printHeader("Channel + WaitGroup")

	ch := make(chan int, len(numbers))
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ch <- n * n
		}(num)
	}

	wg.Wait()
	close(ch)

	sum := 0
	for square := range ch {
		sum += square
	}

	printSum(sum)
}

// Вариант 4: Конвейерная обработка
func sumSquaresPipeline(numbers []int) {
	printHeader("Pipeline")

	numberChan := make(chan int, len(numbers))
	squareChan := make(chan int, len(numbers))

	go func() {
		for _, num := range numbers {
			numberChan <- num
		}
		close(numberChan)
	}()

	go func() {
		for num := range numberChan {
			squareChan <- num * num
		}
		close(squareChan)
	}()

	sum := 0
	for square := range squareChan {
		sum += square
	}

	printSum(sum)
}

// Вариант 5: Используем атомики
func sumSquaresAtomic(numbers []int) {
	printHeader("Atomic Operations")

	var sum int64
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			atomic.AddInt64(&sum, int64(n*n))
		}(num)
	}

	wg.Wait()
	printSum(int(sum))
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	sumSquaresWaitGroup(numbers)
	sumSquaresChannel(numbers)
	sumSquaresChannelWaitGroup(numbers)
	sumSquaresPipeline(numbers)
	sumSquaresAtomic(numbers)
}
