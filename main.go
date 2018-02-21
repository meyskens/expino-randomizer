package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	resty "gopkg.in/resty.v1"
)

var endpoint = os.Getenv("ENDPOINT")

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		doRandom()
		time.Sleep(time.Second)
	}
}

func doRandom() {
	duration, _ := time.ParseDuration("100h")
	data := APIData{
		Setup:     "randomizer",
		Timestamp: time.Now(),
		TTL:       duration,
		Data: []SensorData{
			SensorData{
				Sensor: "x1",
				Data:   rand.Float64(),
			},
			SensorData{
				Sensor: "x2",
				Data:   rand.Float64(),
			},
			SensorData{
				Sensor: "y1",
				Data:   rand.Float64(),
			},
			SensorData{
				Sensor: "y2",
				Data:   rand.Float64(),
			},
		},
	}

	doRequest(data)
}

func doRequest(data APIData) {
	r, _ := resty.R().SetBody(data).Post(endpoint + "/metric")
	fmt.Println(string(r.Body()))
}
