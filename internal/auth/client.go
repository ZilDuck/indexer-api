package auth

import (
	"errors"
	"github.com/google/uuid"
)

var clients []Client

var ErrNoClientFound = errors.New("no auth clients configured")

type Client struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string    `json:"username"`
	ApiKey   string `json:"key"`
	Active   bool      `json:"status"`
}

func GetApiClients() []Client {
	return clients
}

func GetApiClient(apiKey string) (*Client, error) {
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