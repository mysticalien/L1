package main

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func binarySearch(arr []int, elem int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == elem {
			return mid
		} else if arr[mid] < elem {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{2, 4, 6, 8, 9, 12, 18, 24}
	elem := 12

	index := binarySearch(arr, elem)
	if index != -1 {
		fmt.Printf("Element found at index %d\n", index)
	} else {
		fmt.Println("Element not found")
	}
}
