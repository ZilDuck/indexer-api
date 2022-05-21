package auth

import (
	"errors"
)

var clients []Client

var ErrNoClientFound = errors.New("no auth clients configured")

type Client struct {
	ID       uint      `gorm:"primaryKey"`
	Username string    `json:"username"`
	ApiKey   string    `json:"key"`
	Active   bool      `json:"status"`
}

func (c Client) IsAdmin() bool {
	return false
}

func GetApiClients() []Client {
	return clients
}

func GetClientByApiKey(apiKey string) (*Client, error) {
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

func GetClientByUsername(username string) (*Client, error) {
	clients := GetApiClients()

	if len(clients) == 0 {
		return nil, ErrNoClientFound
	}
	for idx := range clients {
		if clients[idx].Username == username {
			return &clients[idx], nil
		}
	}

	return nil, errors.New("username not found")
}