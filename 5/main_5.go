package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	Intersection(a, b)
}

func Intersection(a, b []int) (bool, []int) {
	small, big := a, b
	if len(a) > len(b) {
		small, big = b, a
	}

	set := make(map[int]struct{}, len(small))
	result := make([]int, 0, len(small))

	for _, v := range small {
		set[v] = struct{}{}
	}

	for _, v := range big {
		if _, exists := set[v]; exists {
			result = append(result, v)
		}
	}
	fmt.Printf("%-18s %v\n%-18s %v\n\n",
		"Пересечение:", len(result) > 0,
		"Слайс пересечений:", result,
	)

	return len(result) > 0, result
}
