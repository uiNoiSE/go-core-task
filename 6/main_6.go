package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	randomChan := NewRandomGenerator()

	for i := range 5 {
		num := <-randomChan
		fmt.Printf("Число %d: %d\n", i+1, num)
	}
}

func NewRandomGenerator() <-chan int {
	ch := make(chan int)

	go func() {
		for {
			randomNumber := rand.Int()
			ch <- randomNumber
		}
	}()

	return ch
}
