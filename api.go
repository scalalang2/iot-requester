package main

import (
	"log"
	"net/http"
	"sync"
)

var (
	SERVER_URL = "http://127.0.0.1:3000/"
	REQUEST_URI = "specialEventGet?id=1"
)

func CreateRequest(wg *sync.WaitGroup) {
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