package work

import (
	"log"
	"sync"
)

type Worker interface {
	Task()
}

type WorkerPool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// standard producer-consumer with number of goroutine produce messages and seperate number of threads consume messages
func New(maxGoroutines int) *WorkerPool {
	p := WorkerPool{
		// use unbounded channel to balance the work between different goroutine
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func(id int) {
			defer p.wg.Done()
			for w := range p.work {
				log.Printf("worker %d is running\n", id)
				w.Task()
			}
		}(i)
	}

	return &p
}

func (p *WorkerPool) Submit(w Worker) {
	p.work <- w
}

func (p *WorkerPool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
