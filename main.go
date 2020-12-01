package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	startTime := time.Now().UnixNano()
	fmt.Println("started sending 1000 requests")
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go CreateRequest(&wg)
	}

	fmt.Println("waiting to finish..")
	wg.Wait()

	endTime := time.Now().UnixNano()
	var elapsedTime float64
	elapsedTime = float64(endTime - startTime) / 1e+09
	fmt.Printf("total elapsed time: %.4fs", elapsedTime)
}