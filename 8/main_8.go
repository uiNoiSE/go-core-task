package main

type SemaphoreWaitGroup struct {
	ch    chan struct{}
	count int
}

func (swg *SemaphoreWaitGroup) Add(delta int) {
	swg.ch = make(chan struct{}, delta)
	swg.count = delta
}

func (swg *SemaphoreWaitGroup) Done() {
	swg.ch <- struct{}{}
}

func (swg *SemaphoreWaitGroup) Wait() {
	for range swg.count {
		<-swg.ch
	}
}
