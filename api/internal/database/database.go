package database

import (
	"errors"
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/auth"
	"github.com/ZilDuck/indexer-api/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	ErrorDatabaseConnection = errors.New("failed to connect to database")
)

func NewConnection(cfg config.DB) (*gorm.DB, error) {
	connString := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s&options=%s&sslrootcert=%s",
		cfg.Dialect, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SslMode, cfg.Options, cfg.RootCA)
		zap.L().With(zap.String("conn", connString)).Debug("Database Connection")

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: cfg.Name+"."}})

	if err != nil {
		zap.L().With(zap.Error(err)).Error("error configuring the database")
		err = ErrorDatabaseConnection
		return nil, err
	}

	db.AutoMigrate(&auth.Client{})

	return db, nil
}