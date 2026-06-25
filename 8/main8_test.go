package main

import (
	"testing"
	"time"
)

func TestSemaphoreWaitGroup(t *testing.T) {
	var swg SemaphoreWaitGroup

	numWorkers := 42
	swg.Add(numWorkers)

	workersDone := make(chan bool, numWorkers)

	for i := range numWorkers {
		go func(id int) {
			defer swg.Done()

			time.Sleep(20 * time.Millisecond)
			workersDone <- true
		}(i)
	}

	swg.Wait()

	close(workersDone)

	count := 0
	for range workersDone {
		count++
	}

	if count != numWorkers {
		t.Errorf("Ожидали, что завершится: %d воркеров, но завершилось: %d", numWorkers, count)
	}
}

func TestSemaphoreWaitGroup_Reuse(t *testing.T) {
	var swg SemaphoreWaitGroup

	done := make(chan struct{})

	go func() {
		swg.Add(2)
		go func() { defer swg.Done() }()
		go func() { defer swg.Done() }()
		swg.Wait()

		swg.Add(3)
		go func() { defer swg.Done() }()
		go func() { defer swg.Done() }()
		go func() { defer swg.Done() }()
		swg.Wait()

		close(done)
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("Тест завис. Wait() не разблокировал поток")
	}
}
