package entity

import "time"

type Audit struct {
	Time       time.Time `json:"@timestamp"`
	ApiKey     string    `json:"api-key"`
	Request    string    `json:"request"`
	Network    string    `json:"network"`
	RemoteAddr string    `json:"remote-addr"`
}

