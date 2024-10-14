package main

import (
	"fmt"
	"math/rand"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := rand.Intn(len(arr))
	left, right := 0, len(arr)-1
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[pivot] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	arr := []int{177, -20, 0, 88, -2, 99}
	fmt.Println("Unsorted array:", arr)
	arr = quickSort(arr)
	fmt.Println("Sorted array:", arr)
}
