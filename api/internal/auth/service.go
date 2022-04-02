package auth

import (
	"context"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	LoadClients()
	RefreshClients()
	CreateNewClient(username string, active bool) error
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
			zap.S().Infof("Loaded client %s", loaded[idx].Username)
			clients = append(clients, loaded[idx])
		}
	}
}

func (s service) RefreshClients() {
	var loaded []Client
	s.db.Find(&loaded)

	if len(loaded) == 0 {
		zap.L().Warn("No api keys found in the Database")
	}

	clients = []Client{}

	for idx := range loaded {
		if loaded[idx].Active == true {
			zap.S().Infof("Loaded client %s", loaded[idx].Username)
			clients = append(clients, loaded[idx])
		}
	}
}

func (s service) CreateNewClient(username string, active bool) error {
	return crdbgorm.ExecuteTx(context.Background(), s.db, nil,
		func(tx *gorm.DB) error {
			return s.db.Create(&Client{
				ID:       uuid.New(),
				Username: username,
				ApiKey:   uuid.New().String(),
				Active:   active,
			}).Error
		},
	)
}