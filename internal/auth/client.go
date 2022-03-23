package auth

import (
	"encoding/json"
	"errors"
	"time"
)

var clients []ApiClient

var ErrNoClientFound = errors.New("no auth clients configured")

type ApiClient struct {
	ApiKey   string        `json:"key"`
	Duration time.Duration `json:"duration"`
	Capacity int64         `json:"capacity"`
}

func LoadApiClients(data []string) {
	clients = make([]ApiClient, 0)
	for idx := range data {
		client := ApiClient{}
		if err := json.Unmarshal([]byte(data[idx]), &client); err == nil {
			clients = append(clients, client)
		}
	}
}

func GetApiClients() []ApiClient {
	return clients
}

func GetApiClient(apiKey string) (*ApiClient, error) {
	clients := GetApiClients()

	if len(clients) == 0 {
		return nil, ErrNoClientFound
	}
	for idx := range clients {
		if clients[idx].ApiKey == apiKey {
			return &clients[idx], nil
		}
	}

	return nil, errors.New("API Key not found")
}