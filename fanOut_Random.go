package main

// This fan-out is random type

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func generator[T any, K any](done <-chan K, fn func() T) <-chan T {
	outStream := make(chan T)
	go func() {
		defer close(outStream)
		for {
			select {
			case <-done:
				return
			case outStream <- fn():
			}
		}
	}()
	return outStream
}

func processJob(workerId int, inputNum int) {
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("[Worker %d] job: %d\n", workerId, inputNum)
}

// Worker consumes values and processes them
func worker[T any](id int, done <-chan T, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			processJob(id, job)
		}
	}
}

func main() {
	done := make(chan bool)
	maxLim := 1000000000
	randGen := func() int {
		return rand.Intn(maxLim)
	}

	jobs := generator(done, randGen)

	// Fan-out to multiple workers
	CPUCt := runtime.NumCPU()
	var wg sync.WaitGroup

	for i := 1; i <= CPUCt; i++ {
		wg.Add(1)
		go worker(i, done, jobs, &wg)
	}
	time.Sleep(5 * time.Second)
	close(done)
	wg.Wait()
	fmt.Println("All workers done.")
}
