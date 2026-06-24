package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	Filter(slice1, slice2)
}

func Filter(slice1, slice2 []string) []string {
	set := make(map[string]struct{}, len(slice2))
	result := make([]string, 0, len(slice1))

	for _, v := range slice2 {
		set[v] = struct{}{}
	}

	fmt.Printf("Set:	%v\n\n", set)

	for _, v := range slice1 {
		if _, exists := set[v]; !exists {
			result = append(result, v)
		}
	}

	fmt.Printf("%-12s %v\n%-12s %v\n%-12s %v\n\n",
		"slice1:", slice1,
		"slice2:", slice2,
		"Результат:", result,
	)

	return result
}
