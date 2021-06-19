package main

import (
	"concurrency_pattern/pool"
	"concurrency_pattern/runner"
	"concurrency_pattern/work"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	timeout        = 4 * time.Second
	maxGoroutines  = 25
	pooledResource = 5
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"json",
}

func main() {
	callRunner()
	callDB()
	callWorker()
}

func callRunner() {
	log.Println("Starting work")
	r := runner.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}

	log.Println("Process ended")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

func callDB() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(pool.CreateConnection, pooledResource)

	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			defer wg.Done()
			performQueries(q, p)
		}(query)
	}

	wg.Wait()

	log.Println("Shutdown Program")
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*pool.DbConnection).ID)
}

func callWorker() {
	p := work.New(3)
	var wg sync.WaitGroup
	wg.Add(10 * len(names))
	for i := 0; i < 10; i++ {
		for _, name := range names {
			np := work.NamePrinter{
				Name: name,
			}

			go func() {
				defer wg.Done()
				p.Submit(&np)
			}()
		}
	}

	wg.Wait()
	p.Shutdown()
}
