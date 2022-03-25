package database

import (
	"errors"
	"github.com/ZilDuck/indexer-api/internal/auth"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	ErrorDatabaseConnection = errors.New("failed to connect to database")
)

func NewConnection(connString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "defaultdb.",
		},
	})
	if err != nil {
		zap.L().With(zap.Error(err)).Error("error configuring the database")
		err = ErrorDatabaseConnection
		return nil, err
	}

	db.AutoMigrate(&auth.Client{})

	return db, nil
}