package main

import "fmt"

func main() {
	in := make(chan uint8)
	out := make(chan float64)
	vals := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go CubePipeline(in, out)

	go func() {
		for _, v := range vals {
			in <- v
		}
		close(in)
	}()

	for res := range out {
		fmt.Println(res)
	}
}

func CubePipeline(in <-chan uint8, out chan<- float64) {
	for num := range in {
		res := float64(num)
		out <- res * res * res
	}
	close(out)
}
