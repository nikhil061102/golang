package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TEE pattern is when one values from one input channel are distributed into n channels
// FAN-OUT pattern when values go into channels as per availabilty of the channels at that time
//        input_chan => [outChan1, outChan2, outChan3]
// TEE:        A     => [   A    ,    A    ,    A    ]
// FAN-OUT:    A     => [   -    ,    A    ,    -    ]

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

func roundRobinFanOut[K any, T any](done <-chan K, inputStream <-chan T, numOutChans int) []chan T {
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

		idx := 0 // round-robin index
		for {
			select {
			case <-done:
				return
			case val, ok := <-inputStream:
				if !ok {
					return
				}
				select {
				case <-done:
					return
				case outChanSlice[idx] <- val:
					idx = (idx + 1) % numOutChans // increment index
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
	outs := roundRobinFanOut(done, randStream, 3)

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
	time.Sleep(2 * time.Second)
	close(done)
}
