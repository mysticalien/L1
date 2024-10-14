package main

import (
	"fmt"
	"sort"
)

/*
Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func groupTemperatures(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64)
	for _, temp := range temps {
		key := int(temp/10) * 10
		groups[key] = append(groups[key], temp)
	}
	return groups
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := groupTemperatures(temps)

	var keys []int
	for key := range groups {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	fmt.Println("Grouped temperatures:")
	for _, key := range keys {
		fmt.Printf("%d: %v\n", key, groups[key])
	}
}
