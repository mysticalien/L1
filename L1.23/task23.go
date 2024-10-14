package main

import (
	"fmt"
	"log"
)

/*
Удалить i-ый элемент из слайса
*/

// Удаление i-го элемента из слайса без сохранения порядка O(1)
func removeElement(arr []int, i int) ([]int, error) {
	if i < 0 || i >= len(arr) {
		return nil, fmt.Errorf("index %d out of bounds", i)
	}
	arr[i] = arr[len(arr)-1] // Копируем последний элемент на место i
	return arr[:len(arr)-1], nil
}

// Удаление i-го элемента из слайса с сохранением порядка O(n)
func removeElementOrdered(arr []int, i int) ([]int, error) {
	if i < 0 || i >= len(arr) {
		return nil, fmt.Errorf("index %d out of bounds", i)
	}
	return append(arr[:i], arr[i+1:]...), nil
}

// Удаление i-го элемента с сдвигом влево с сохранением порядка O(n)
func shiftLeft(arr []int, i int) ([]int, error) {
	if i < 0 || i >= len(arr) {
		return nil, fmt.Errorf("index %d out of bounds", i)
	}
	copy(arr[i:], arr[i+1:]) // Сдвиг влево
	return arr[:len(arr)-1], nil
}

func main() {
	fmt.Println("Remove Element")
	arr := []int{1, 2, 3, 4, 5}
	var err error
	arr, err = removeElement(arr, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arr)

	fmt.Println("Remove Element Ordered")
	arr1 := []int{1, 2, 3, 4, 5}
	arr1, err = removeElementOrdered(arr1, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arr1)

	fmt.Println("Shift Left")
	arr2 := []int{1, 2, 3, 4, 5}
	arr2, err = shiftLeft(arr2, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arr2)
}
