package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров
*/

/* Обоснование выбора: Для завершения работы воркеров выбран метод закрытия канала.
Этот подход проще и интуитивнее:
	воркеры автоматически завершаются при попытке прочитать из закрытого канала.
	В отличие от использования контекста, который требует дополнительного управления,
	закрытие канала упрощает реализацию и улучшает читаемость кода, избегая лишних проверок состояния.
*/

func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Worker %d: received %d\n", id, data)
	}
	fmt.Printf("Worker %d: stopped\n", id)
}

func produceData(ch chan<- int) {
	go func() {
		i := 0
		for {
			ch <- i
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

func main() {
	var workerCount int
	fmt.Print("Enter the number of workers: ")
	_, err := fmt.Scan(&workerCount)
	if err != nil || workerCount <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}

	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	produceData(ch)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh

	close(ch)
	wg.Wait()
	fmt.Println("All workers stopped")
}
