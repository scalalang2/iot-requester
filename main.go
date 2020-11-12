package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	SERVER_URL = "http://127.0.0.1:3000/"
	REQUEST_URI = "specialEventGet?id=1"
)

func createRequest(wg *sync.WaitGroup) {
	defer wg.Done()
	url := SERVER_URL + REQUEST_URI

	resp, err := http.Get(url)

	if resp.StatusCode != 200 {
		log.Println("the status code received is not ok")
	}

	if err != nil {
		log.Println("http client error!")
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func main() {
	var wg sync.WaitGroup

	startTime := time.Now().UnixNano()

	fmt.Println("started sending 1000 requests")
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go createRequest(&wg)
	}

	fmt.Println("waiting to finish..")
	wg.Wait()

	endTime := time.Now().UnixNano()
	var elapsedTime float64
	elapsedTime = float64(endTime - startTime) / 1e+09
	fmt.Printf("total elapsed time: %.4fs", elapsedTime)
}