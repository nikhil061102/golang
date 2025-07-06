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

func tee[K any, T any](done <-chan K, inputStream <-chan T, numOutChans int) []chan T {
	outChanSlice := make([]chan T, numOutChans)
	for i := 0; i < numOutChans; i++ {
		ch := make(chan T)
		outChanSlice[i] = ch
	}

	go func() {
		defer func() {
			for _, ch := range outChanSlice {
				ch := ch
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
				for _, ch := range outChanSlice {
					ch := ch
					ch <- val
				}
			}
		}
	}()

	return outChanSlice
}

func main() {
	done := make(chan bool)
	maxLim := 100

	randGen := func() int {
		return rand.Intn(maxLim)
	}

	transactionsStream := generator(done, randGen)
	outs := tee(done, transactionsStream, 3)

	names := []string{"Analytics", "Fraud", "Finance"}

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
