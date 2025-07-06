// Use confinement to avoid mutex locks (bcoz unnecessary overhead)

// 1. race condition without locks in writing to array
// 2. doubleInt in the lock. But no need since it can not possibly cause race cond in writing
// 3. So removed it from lock & time reduced from 5 secs to 1 sec
// 4. Confined array uses indexing & gives same ordered output & that too without mutex locks 

package main

import (
	"fmt"
	"sync"
	"time"
)

func doubleInt (num int) int {
	time.Sleep(time.Second)
	return num * 2;
}

func updateResArray (res *[]int, num int, wg *sync.WaitGroup, mu *sync.Mutex){
	defer wg.Done()

	processedNum := doubleInt(num)
	mu.Lock()
	// processedNum := doubleInt(num) 
	*res = append(*res, processedNum)
	mu.Unlock()
}

func confinementUpdate(res *int, num int, wg *sync.WaitGroup){
	defer wg.Done()
	processedNum := doubleInt(num)

	*res = processedNum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	// res := []int{}
	confinedSliceSize := len(nums)
	res := make([]int, confinedSliceSize)
	now := time.Now()
	
	// var mu sync.Mutex
	var wg sync.WaitGroup
	
	for idx, num := range nums {
		wg.Add(1)
		// go updateResArray(&res, num, &wg, &mu)
		go confinementUpdate(&res[idx], num, &wg)
	}

	wg.Wait()
	fmt.Println("Time taken := ", time.Since(now))

	fmt.Println("Doubled numbers:", res)
}
