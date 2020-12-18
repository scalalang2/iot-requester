package api

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var (
	SERVER_URL = "<SERVER_URI>"
	Params = make([]*ReqParams, 20000)
	Number = 0
	Alerts = []string{"warning", "normal", "danger"}
)

type ReqParams struct {
	Id string
	KeyValue string
	DeviceId string
	SensorId string
	SensorCategoryId string
	SensorValue string
	SensorAlertMsg string
	SensorDescription string
	EventCreationTime string
}

func PrepareReqBody() {
	for i := 0; i < 20000; i++ {
		localParams := ReqParams {
			fmt.Sprintf("ID_%d", i+1),
			strconv.Itoa(rand.Intn(10)),
			strconv.Itoa(rand.Intn(100)),
			strconv.Itoa(rand.Intn(6)),
			strconv.Itoa(rand.Intn(3)),
			strconv.Itoa(rand.Intn(5) + 20),
			Alerts[rand.Intn(3)],
			"normal",
			strconv.Itoa(rand.Intn(10) + 15),
		}
		Params[i] = &localParams
	}
}

func createParams() *ReqParams {
	mux := &sync.Mutex{}
	mux.Lock()

	localParams := Params[Number]

	Number++
	mux.Unlock()
	return localParams
}

func CreateRequest(requests chan bool, queues chan bool) {
	params := createParams()

	url := fmt.Sprintf("%slocationSet?id=%s&keyValue=%s&deviceId=%s&sensorId=%s&sensorCategoryId=%s&sensorValue=%s&sensorAlertMsg=%s&sensorDescription=%s&eventCreateTime=%s",
		SERVER_URL,
		params.Id,
		params.KeyValue,
		params.DeviceId,
		params.SensorId,
		params.SensorCategoryId,
		params.SensorValue,
		params.SensorAlertMsg,
		params.SensorDescription,
		params.EventCreationTime)

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
