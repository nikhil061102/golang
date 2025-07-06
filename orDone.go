package main

import (
	"fmt"
	"sync"
)

func generator(out chan<- string) {
	defer close(out)
	for i := 1; i <= 100; i++ {
		out <- fmt.Sprintf("Value-%d", i)
	}
}

func doWorkFirst(done <-chan bool, in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		case val, ok := <-in:
			if !ok {
				return
			}
			fmt.Printf("1. Input coming is %s\n", val)
		}
	}
}

func doWorkSecond(done <-chan bool, in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		case val, ok := <-in:
			if !ok {
				return
			}
			fmt.Printf("2. Input coming is %s\n", val)
		}
	}
}

func doWorkThird(done <-chan bool, in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		case val, ok := <-in:
			if !ok {
				return
			}
			fmt.Printf("3. Input coming is %s\n", val)
		}
	}
}

// now we want to extract the for select done code as separate so that no code duplication
func orDone(done <-chan bool, input <-chan string) <-chan string {
	relayStream := make(chan string)

	go func() {
		defer close(relayStream)
		for {
			select {
			case <-done:
				return
			case val, ok := <-input:
				if !ok {
					return
				}
				select {
				case relayStream <- val:
				case <-done:
					return
				}

			}
		}
	}()

	return relayStream
}

func doWorkFirstImproved(done <-chan bool, in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch := orDone(done, in)
	for {
		val, ok := <-ch 
		if !ok {
			return
		}
		fmt.Printf("1. Input coming is %s\n", val)
	}
}

func doWorkSecondImproved(done <-chan bool, in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch := orDone(done, in)
	for {
		val, ok := <-ch 
		if !ok {
			return
		}
		fmt.Printf("2. Input coming is %s\n", val)
	}
}

func doWorkThirdImproved(done <-chan bool, in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch := orDone(done, in)
	for {
		val, ok := <-ch 
		if !ok {
			return
		}
		fmt.Printf("3. Input coming is %s\n", val)
	}
}

func main() {
	in := make(chan string)
	done := make(chan bool)
	var wg sync.WaitGroup

	go generator(in)
	wg.Add(3)
	// go doWorkFirst(done, in, &wg)
	// go doWorkSecond(done, in, &wg)
	// go doWorkThird(done, in, &wg)
	go doWorkFirstImproved(done, in, &wg)
	go doWorkSecondImproved(done, in, &wg)
	go doWorkThirdImproved(done, in, &wg)

	wg.Wait()
	fmt.Println("Hello")
}
