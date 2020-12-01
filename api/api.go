package api

import (
	"fmt"
	"log"
	"net/http"
)

var (
	SERVER_URL = "http://220.76.205.230:3000/"
)

type ReqParams struct {
	id string
	keyValue string
	deviceId string
	sensorId string
	sensorCategoryId string
	sensorValue string
	sensorAlertMsg string
	sensorDescription string
	eventCreationTime string
}

func createParams() *ReqParams {
	params := ReqParams {
		"idTesting2",
		"1",
		"3",
		"5",
		"2",
		"12",
		"warning",
		"error",
		"22",
	}
	return &params
}

func CreateRequest(requests chan bool, queues chan bool) {
	params := createParams()

	url := fmt.Sprintf("%slocationSet?id=%s&keyValue=%s&deviceId=%s&sensorId=%s&sensorCategoryId=%s&sensorValue=%s&sensorAlertMsg=%s&sensorDescription=%s&eventCreateTime=%s",
		SERVER_URL,
		params.id,
		params.keyValue,
		params.deviceId,
		params.sensorId,
		params.sensorCategoryId,
		params.sensorValue,
		params.sensorAlertMsg,
		params.sensorDescription,
		params.eventCreationTime)

	resp, err := http.Post(url, "application/json", nil)

	if resp != nil && resp.StatusCode != 200 {
		//log.Println("the status code received is not ok")
	}

	if err != nil {
		log.Println("http client error!")
		log.Fatal(err)
	}

	err = resp.Body.Close()
	if err != nil {
		log.Println("body close error!")
	}

	requests <- true
	queues <- true
}