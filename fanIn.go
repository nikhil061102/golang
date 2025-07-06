package main

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

func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()
	return taken
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= num-1; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primesGenerator[K any](done <-chan K, inputStream <-chan int) <-chan int {
	primeStream := make(chan int)
	go func() {
		defer close(primeStream)
		for {
			select {
			case <-done:
				return
			case num, ok := <-inputStream:
				if !ok {
					return
				}
				if isPrime(num) {
					select {
					case <-done:
						return
					case primeStream <- num:
					}
				}
			}
		}
	}()
	return primeStream
}

func fanIn[K any, T any](done <-chan K, inputChans ...<-chan T) <-chan T {
	resChan := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(inputChans))

	for _, ch := range inputChans {
		ch := ch // 👈 capture correct reference (Check Loop-Variable-Capture.md)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				case val, ok := <-ch:
					if !ok {
						return
					}
					select {
					case <-done:
						return
					case resChan <- val:
					}
				}
			}
		}()
	}

	// Close resChan when all goroutines are done
	go func() {
		wg.Wait()
		close(resChan)
	}()

	return resChan
}

func main() {
	done := make(chan bool)
	maxLim := 1000000000
	randGen := func() int {
		return rand.Intn(maxLim)
	}
	now := time.Now()
	outChan := generator(done, randGen)
	// primeChan := primesGenerator(done, outChan)

	// fan out = for each CPU we will spin up a different goroutine that are then combined into one chan for returning primes
	CPUCt := runtime.NumCPU()
	// fmt.Println("No. of CPUs =", CPUCt)
	primeFindingChannels := make([]<-chan int, CPUCt)
	for i := 0; i < CPUCt; i++ {
		primeFindingChannels[i] = primesGenerator(done, outChan)
	}
	combinedChan := fanIn(done, primeFindingChannels...)
	limitedOut := take(done, combinedChan, 20)

	for {
		val, ok := <-limitedOut
		if !ok {
			break
		}
		fmt.Println(val)
	}
	fmt.Println(time.Since(now))
	// Time taken in pipeline type = 32 secs (approx) & in fanIn-Out = 5 sec
	close(done)
	fmt.Println("khatam sab")
}
