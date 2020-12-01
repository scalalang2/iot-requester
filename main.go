package main

import (
	"fmt"
	"github.com/fatih/color"
	"iot-requester/api"
	"net/http"
	"time"
)

func serve(requests chan bool, ticker <-chan time.Time, done chan bool) {
	sum := 0
	count := 1
	i := 0
	for {
		select {
			case <-requests:
				sum++
				i++
			case time := <-ticker:
				reportText := color.New(color.FgBlack).Add(color.BgGreen).PrintfFunc()
				pointText := color.New(color.FgRed).PrintfFunc()

				fmt.Printf("%s\n", time)
				reportText("[Report]")
				fmt.Printf(" Served ")
				pointText("%d", i)
				fmt.Printf(" Requests, TPS: ")
				pointText("%.2f\n\n", float64(sum)/float64(count))
				count++
				i = 0

				if count == 21 {
					done<-true
				}
		}
	}
}

func makeRequests(requests chan bool, queues chan bool) {
	limiter := time.Tick(1500 * time.Microsecond)

	for {
		<-limiter
		<-queues
		go api.CreateRequest(requests, queues);
	}
}

func main() {
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000

	done := make(chan bool)
	ticker := time.NewTicker(time.Second)
	requests := make(chan bool, 1000)
	workQueue := make(chan bool, 1000)
	for i := 0; i < 1000; i++ {
		workQueue<-true;
	}

	fmt.Println("Start to make requests")
	go serve(requests, ticker.C, done);
	go makeRequests(requests, workQueue);

	<-done
}