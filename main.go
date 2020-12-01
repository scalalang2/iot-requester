package main

import (
	"encoding/csv"
	"fmt"
	"github.com/fatih/color"
	"iot-requester/api"
	"log"
	"net/http"
	"os"
	"time"
)

func serve(requests chan bool, ticker <-chan time.Time, done chan bool) {
	sum := 0
	count := -2
	i := 0
	for {
		select {
			case <-requests:
				if count >= 1 {
					sum++
				}
				i++
			case tickTime := <-ticker:
				if count >= 1 {
					reportText := color.New(color.FgBlack).Add(color.BgGreen).PrintfFunc()
					pointText := color.New(color.FgRed).PrintfFunc()

					fmt.Printf("%s\n", tickTime.Format(time.UnixDate))
					reportText("[Report]")
					fmt.Printf(" Served ")
					pointText("%d", i)
					fmt.Printf(" Requests, TPS: ")
					pointText("%.2f\n\n", float64(sum)/float64(count))
				}

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

func reportCSV() {
	params := api.Params
	file, err := os.Create("input.csv")
	if err != nil {
		log.Println("Cannot report data to csv file.")
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 20000; i++ {
		item := params[i]
		if item == nil {
			break;
		}

		err = writer.Write([]string{
			item.Id,
			item.KeyValue,
			item.DeviceId,
			item.SensorId,
			item.SensorCategoryId,
			item.SensorValue,
			item.SensorAlertMsg,
			item.SensorDescription,
			item.EventCreationTime,
		})

		if err != nil {
			log.Println("Cannot report data to csv file.")
			log.Fatal(err)
		}
	}
}

func main() {
	api.PrepareReqBody()
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 2000

	done := make(chan bool)
	ticker := time.NewTicker(time.Second)
	requests := make(chan bool, 1000)
	workQueue := make(chan bool, 1000)

	for i := 0; i < 1000; i++ {
		workQueue<-true
	}

	fmt.Println("Start to make requests")
	go serve(requests, ticker.C, done);
	go makeRequests(requests, workQueue);

	<-done
	reportCSV()
	fmt.Println()

	greenColor := color.New(color.FgGreen).PrintlnFunc()
	greenColor("Saved input.csv file.")
}