package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func intersectSets(set1, set2 []int) []int {
	setMap := make(map[int]struct{})
	result := []int{}

	for _, value := range set1 {
		setMap[value] = struct{}{}
	}

	for _, value := range set2 {
		if _, exists := setMap[value]; exists {
			result = append(result, value)
		}
	}

	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	set2 := []int{5, 4, 3, 8, 10, 11, 12, 1}

	result := intersectSets(set1, set2)
	fmt.Println("Intersection:", result)
}
