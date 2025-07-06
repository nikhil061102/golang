package main

import (
	"fmt"
	"math/rand"
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

func main() {
	done := make(chan bool)
	maxLim := 100000
	randGen := func() int {
		return rand.Intn(maxLim)
	}
	
	outChan := generator(done, randGen)
	limitedOut := take(done, outChan, 20)
	for {
		val, ok := <-limitedOut
		if !ok {
			break 
		}
		fmt.Println(val)
	}
	close(done)
}
