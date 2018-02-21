package main

import "time"

// APIData is the data sent by an API call
type APIData struct {
	Setup     string        `json:"setup"`
	Timestamp time.Time     `json:"timestamp"`
	TTL       time.Duration `json:"ttl"`
	Data      []SensorData  `json:"data"`
}

// SensorData is the data of one sensor sent an an APIData
type SensorData struct {
	Sensor string  `json:"sensor"`
	Data   float64 `json:"data"`
}
