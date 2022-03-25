package entity

type Audit struct {
	ApiKey     string `json:"apiKey"`
	Request    string `json:"request"`
	Network    string `json:"network"`
	RemoteAddr string `json:"remoteAddr"`
}

