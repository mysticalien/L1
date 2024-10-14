package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func createSet(strings []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, str := range strings {
		set[str] = struct{}{}
	}
	return set
}

func setToSlice(set map[string]struct{}) []string {
	slice := make([]string, 0, len(set))
	for str := range set {
		slice = append(slice, str)
	}
	return slice
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	set := createSet(strings)
	setSlice := setToSlice(set)

	fmt.Println("Set from strings:", setSlice)
}
