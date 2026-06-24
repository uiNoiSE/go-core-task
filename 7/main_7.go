package main

import "sync"

func main() {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	ch1 <- 1
	ch1 <- 2
	ch2 <- 3
	ch2 <- 4

	close(ch1)
	close(ch2)

	merged := Merge(ch1, ch2)
	for val := range merged {
		println("Получено:", val)
	}
}

func Merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	wg.Add(len(channels))

	for _, ch := range channels {
		go func(ch <-chan int) {
			defer wg.Done()

			for val := range ch {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
