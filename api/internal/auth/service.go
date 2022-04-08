package auth

import (
	"context"
	"errors"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	LoadClients()
	RefreshClients()
	CreateClient(username string, active bool) (*Client, error)
	UpdateClient(client Client) error
	DeleteClient(client Client) error
}

type service struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) Service {
	return service{db: db}
}

func (s service) LoadClients() {
	var loaded []Client
	s.db.Find(&loaded)

	if len(loaded) == 0 {
		zap.L().Warn("No api keys found in the Database")
	}

	for idx := range loaded {
		if loaded[idx].Active == true {
			zap.S().With(zap.String("apikey", loaded[idx].ApiKey)).Infof("Loaded client %s-%s", loaded[idx].Username, loaded[idx].ID)
			clients = append(clients, loaded[idx])
		}
	}
}

func (s service) RefreshClients() {
	zap.L().Info("Refreshing clients")
	var loaded []Client
	s.db.Find(&loaded)

	if len(loaded) == 0 {
		zap.L().Warn("No api keys found in the Database")
	}

	clients = []Client{}

	for idx := range loaded {
		if loaded[idx].Active == true {
			clients = append(clients, loaded[idx])
		}
	}
}

func (s service) CreateClient(username string, active bool) (*Client, error) {
	zap.S().Infof("Create client %s", username)

	s.RefreshClients()

	for idx := range clients {
		if clients[idx].Username == username {
			return nil, errors.New("username already in use")
		}
	}

	client := &Client{
		ID:       uuid.New(),
		Username: username,
		ApiKey:   uuid.New().String(),
		Active:   active,
	}

	err := crdbgorm.ExecuteTx(context.Background(), s.db, nil,
		func(tx *gorm.DB) error {
			return s.db.Create(client).Error
		},
	)

	if err == nil {
		s.RefreshClients()
	}

	return client, err
}

func (s service) UpdateClient(client Client) error {
	zap.S().Infof("Update client %s", client.Username)

	if err := s.db.Save(&client).Error; err != nil {
		return err
	}

	s.RefreshClients()

	return nil
}

func (s service) DeleteClient(client Client) error {
	zap.S().Infof("Delete client %s", client.Username)

	if err := s.db.Delete(client).Error; err != nil {
		return err
	}

	s.RefreshClients()

	return nil
}