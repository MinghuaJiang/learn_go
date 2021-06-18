package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoRoutine = 8
	taskLoad        = 10
)

func startWorker() {
	var wg sync.WaitGroup
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoRoutine)

	for gr := 1; gr <= numberGoRoutine; gr++ {
		go worker(tasks, gr, &wg)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
