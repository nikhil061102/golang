package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(done <-chan bool, workerId int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Printf("Worker no. %d Do some work... %d\n", workerId, i)
			i++
		}
	}
}

func main() {
	var wg sync.WaitGroup

	done1 := make(chan bool)
	done2 := make(chan bool)
	done3 := make(chan bool)
	wg.Add(3)
	go doWork(done1, 1, &wg)
	go doWork(done2, 2, &wg)
	go doWork(done3, 3, &wg)
	// basically when done sends true then the goroutines stops & returns
	// so parent goroutine (main) controls the children goroutines with done channel
	// if a common channel is sent to all 3 of them, then all 3 stop at the same time.
	time.Sleep(1 * time.Second)
	done1 <- true
	time.Sleep(1 * time.Second)
	done2 <- true
	time.Sleep(1 * time.Second)
	done3 <- true

	wg.Wait()
	fmt.Println("Hello")
}
