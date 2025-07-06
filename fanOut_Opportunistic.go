package main

// This fan-out is opportunistic - value is sent to the first available output (non-blocking)

import (
	"fmt"
	"math/rand"
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

func opportunisticFanOut[K any, T any](done <-chan K, inputStream <-chan T, numOutChans int) []chan T {
	outChanSlice := make([]chan T, numOutChans)
	for i := 0; i < numOutChans; i++ {
		outChanSlice[i] = make(chan T)
	}

	go func() {
		defer func() {
			for _, ch := range outChanSlice {
				close(ch)
			}
		}()

		for {
			select {
			case <-done:
				return
			case val, ok := <-inputStream:
				if !ok {
					return
				}

				delivered := false
				for !delivered {
					select {
					case <-done:
						return
					default:
						for _, ch := range outChanSlice {
							select {
							case ch <- val:
								delivered = true
								break
							default:
								// not ready, try next
							}
						}
					}
				}
			}
		}
	}()

	return outChanSlice
}

func main() {
	done := make(chan bool)
	maxLim := 100000

	randGen := func() int {
		return rand.Intn(maxLim)
	}

	randStream := generator(done, randGen)
	outs := opportunisticFanOut(done, randStream, 3)

	names := []string{"worker1", "worker2", "worker3"}

	// Start 3 receivers
	for i, ch := range outs {
		go func(name string, ch <-chan int) {
			for val := range ch {
				fmt.Printf("[%s] Received: %d\n", name, val)
			}
		}(names[i], ch)
	}

	// Let it run for a short time then stop
	time.Sleep(1 * time.Second)
	close(done)
}
