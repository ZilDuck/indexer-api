package dto

import "time"

type Audit struct {
	Time       time.Time `json:"timestamp"`
	Request    string    `json:"request"`
	Network    string    `json:"network"`
	RemoteAddr string    `json:"remoteAddr"`
}

