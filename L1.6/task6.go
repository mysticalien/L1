package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// Вариант 1: Остановка с использованием context.Context
func workerWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker (context) stopped")
			return
		default:
			fmt.Println("Worker (context) working")
			time.Sleep(1 * time.Second)
		}
	}
}

// Вариант 2: Остановка через закрытие канала
func workerWithChannel(stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			fmt.Println("Worker (channel) stopped")
			return
		default:
			fmt.Println("Worker (channel) working")
			time.Sleep(1 * time.Second)
		}
	}
}

// Вариант 3: Остановка с использованием атомарного флага
func workerWithAtomicFlag(stop *int32) {
	for atomic.LoadInt32(stop) == 0 {
		fmt.Println("Worker (atomic flag) working")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Worker (atomic flag) stopped")
}

// Вариант 4: Остановка с помощью runtime.Goexit()
func workerWithGoexit(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		fmt.Println("Worker (Goexit) working")
		time.Sleep(1 * time.Second)
		runtime.Goexit() // Принудительно завершаем горутину
	}
}

func main() {
	var wg sync.WaitGroup

	// Вариант 1: Использование context.Context для остановки
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go workerWithContext(ctx)

	// Вариант 2: Закрытие канала
	stopCh := make(chan struct{})
	go workerWithChannel(stopCh)

	// Вариант 3: Атомарный флаг
	var stopFlag int32 = 0
	go workerWithAtomicFlag(&stopFlag)

	// Вариант 4: Использование runtime.Goexit()
	wg.Add(1)
	go workerWithGoexit(&wg)

	time.Sleep(3 * time.Second)

	// Закрываем канал для завершения горутины через канал
	close(stopCh)

	// Устанавливаем атомарный флаг для остановки третьей горутины
	atomic.StoreInt32(&stopFlag, 1)

	// Ждем завершения горутины с Goexit
	wg.Wait()

	time.Sleep(1 * time.Second)
	fmt.Println("All workers stopped")
}
