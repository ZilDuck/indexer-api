package entity

import "time"

type Audit struct {
	Time       time.Time `json:"@timestamp"`
	ApiKey     string    `json:"apiKey"`
	Request    string    `json:"request"`
	Network    string    `json:"network"`
	RemoteAddr string    `json:"remoteAddr"`
}
