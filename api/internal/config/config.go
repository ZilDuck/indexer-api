package config

import (
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/log"
	"github.com/spf13/viper"
	"os"
)

var config Config

type Config struct {
	Env           string
	Port          int
	Debug         bool
	ElasticSearch ElasticSearch
	DB            DB
	Aws           Aws
	CdnHost       string
	AuditDir      string
	AdminIds      []string
}

type Aws struct {
	AccessKey string
	SecretKey string
	Region    string
}

type ElasticSearch struct {
	Hosts       []string
	Sniff       bool
	HealthCheck bool
	Debug       bool
	Username    string
	Password    string
}

type DB struct {
	Dialect  string
	Username string
	Password string
	Host     string
	Port     int
	Name     string
	SslMode  string
	Options  string
	RootCA   string
}

func Init() {
	viper.SetConfigName("env.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic("Failed to unmarshal config")
	}

	_ = os.Setenv("AWS_ACCESS_KEY_ID", config.Aws.AccessKey)
	_ = os.Setenv("AWS_SECRET_KEY_ID", config.Aws.SecretKey)
	_ = os.Setenv("AWS_REGION", config.Aws.Region)

	log.NewLogger(config.Debug)
}

func Get() Config {
	return config
}
