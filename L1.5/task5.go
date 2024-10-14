package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

func getProgramDuration() int {
	var n int
	fmt.Println("Enter the program lifetime in seconds:")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return -1
	}
	return n
}

func sender(programDuration int, ch chan int) {
	timer := time.NewTimer(time.Duration(programDuration) * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Println("Timer expired!")
			close(ch)
			return
		default:
			ch <- rand.Intn(100)
			time.Sleep(time.Second)
		}
	}
}

func receiver(ch chan int) {
	for value := range ch {
		fmt.Println("Received:", value)
	}
}

func main() {
	programDuration := getProgramDuration()
	if programDuration == -1 {
		return
	}

	ch := make(chan int)
	go sender(programDuration, ch)

	receiver(ch)
}
