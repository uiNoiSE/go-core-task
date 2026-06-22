package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	originalSlice := make([]int, 10)

	for idx := range originalSlice {
		originalSlice[idx] = rand.IntN(100)
	}
	fmt.Printf("original:	%v,	len:	%v\n", originalSlice, len(originalSlice))

	sliced := SliceExample(originalSlice)
	fmt.Printf("sliced(no odds):	%v,	len:	%v\n", sliced, len(sliced))

	added := AddElements(sliced, len(sliced))
	fmt.Printf("added(sliced + sliced's len):	%v,	len:	%v\n", added, len(added))
	fmt.Printf("CHECK sliced:	%v,	len:	%v\n", sliced, len(sliced))

	copied := CopySlice(added)
	fmt.Printf("copied(copy of added):	%v,	len:	%v\n", copied, len(copied))

	added = append(added, 1010)
	fmt.Printf("CHANGE added:	%v,	len:	%v\n", added, len(added))
	fmt.Printf("CHECK copied:	%v,	len:	%v\n", copied, len(copied))

	removed := RemoveElement(added, 1)
	fmt.Printf("removed(2nd idx from added):	%v,	len: %v\n", removed, len(removed))
	fmt.Printf("CHANGE added:	%v,	len:	%v\n", added, len(added))
}

func SliceExample(s []int) []int {
	var result []int

	for _, v := range s {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	fmt.Printf("Sliced: %v\n", result)

	return result
}

func AddElements(s []int, target int) []int {
	return append(s, target)
}

func CopySlice(s []int) []int {
	newCopy := make([]int, len(s))
	copy(newCopy, s)

	return newCopy
}

func RemoveElement(s []int, idx int) []int {
	local := CopySlice(s)

	result := append(local[:idx], local[idx+1:]...)
	return result
}
