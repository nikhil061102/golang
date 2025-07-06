package main

import (
	"fmt"
	"math"
)

func ch(arr []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, val := range arr {
			out <- val
		}
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			val, ok := <-in
			if !ok {
				return
			}
			out <- val * val
		}
	}()
	return out
}

func add(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			val, ok := <-in
			if !ok {
				return
			}
			out <- val + 10
		}
	}()
	return out
}

func sqrt(in <-chan int) <-chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)
		for {
			val, ok := <-in
			if !ok {
				return
			}
			out <- math.Sqrt(float64(val))
		}
	}()
	return out
}

func generator() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < 100; i++ {
			out <- i
		}
	}()
	return out
}

// we have a array of numbers & we have to find sqrt(x^2 + 10)
func main() {
	// secondStep := sq(generator())
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	firstStep := ch(nums)
	secondStep := sq(firstStep)
	thirdStep := add(secondStep)
	fourthStep := sqrt(thirdStep)

	for {
		val, ok := <-fourthStep
		if !ok {
			return
		}
		fmt.Println(val)
	}
	// now this can also be used with a generator
	
	// secondStep := sq(generator())
	// thirdStep := add(secondStep)
	// fourthStep := sqrt(thirdStep)
}
