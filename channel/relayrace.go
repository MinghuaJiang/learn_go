package main

import (
	"fmt"
	"sync"
	"time"
)

func startRun() {
	var wg sync.WaitGroup
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton, &wg)

	baton <- 1

	wg.Wait()
}

func Runner(baton chan int, wg *sync.WaitGroup) {
	var newRunner int
	runner := <-baton
	fmt.Printf("Runner %d Running With Baton\n", runner)
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To the line \n", newRunner)
		go Runner(baton, wg)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner
}
