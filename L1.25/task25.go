package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

// Сон с использованием time.After
func sleepWithAfter(d time.Duration) {
	<-time.After(d)
}

// Сон с использованием цикла
// Это не самый эффективный способ, так как он постоянно использует CPU в процессе ожидания,
// что может замедлить работу программы
func sleepWithLoop(d time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) >= d {
			return
		}
	}
}

// Сон с использованием таймера
func sleepWithTimer(ns time.Duration) {
	if ns <= 0 {
		return
	}
	t := time.NewTimer(ns)
	<-t.C
}

func main() {
	var duration int
	fmt.Println("Enter number of seconds: ")
	_, err := fmt.Scan(&duration)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Sleeping for %d seconds (Using After)\n", duration)
	sleepWithAfter(time.Duration(duration) * time.Second)
	fmt.Println("Awake!!!")

	fmt.Printf("Sleeping for %d seconds (Using Loop)\n", duration)
	sleepWithLoop(time.Duration(duration) * time.Second)
	fmt.Println("Awake!!!")

	fmt.Printf("Sleeping for %d seconds (Using Timer)\n", duration)
	sleepWithTimer(time.Duration(duration) * time.Second)
	fmt.Println("Awake!!!")
}
