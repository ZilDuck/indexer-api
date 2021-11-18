package config

import (
	"github.com/ZilDuck/indexer-api/internal/log"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Env           string
	Port          int
	Network       string
	Index         string
	Debug         bool
	ElasticSearch ElasticSearchConfig
	Aws           AwsConfig
	SentryDsn     string
}

type AwsConfig struct {
	AccessKey string
	SecretKey string
	Token     string
	Region    string
}

type ElasticSearchConfig struct {
	Aws         bool
	Host        string
	Sniff       bool
	HealthCheck bool
	Debug       bool
	Username    string
	Password    string
}

type ThrottleConfig struct {
	MaxEventsPerSec int
	MaxBurstSize    int
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	initLogger()

	initSentry()
}
func initLogger() {
	log.NewLogger(Get().Debug, Get().SentryDsn)
}

func initSentry() {
	if Get().SentryDsn != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:         Get().SentryDsn,
			Environment: Get().Env,
			Debug:       Get().Debug,
		}); err != nil {
			zap.L().With(zap.Error(err)).Fatal("Sentry init")
		}
	}
}

func Get() *Config {
	return &Config{
		Env:       getString("ENV", ""),
		Port:      getInt("PORT", 8080),
		Network:   getString("NETWORK", "zilliqa"),
		Index:     getString("INDEX_NAME", "xxx"),
		Debug:     getBool("DEBUG", false),
		SentryDsn: getString("SENTRY_DSN", ""),
		Aws: AwsConfig{
			AccessKey: getString("AWS_ES_ACCESS_KEY", ""),
			SecretKey: getString("AWS_ES_SECRET_KEY", ""),
			Token:     getString("AWS_TOKEN", ""),
			Region:    getString("AWS_REGION", ""),
		},
		ElasticSearch: ElasticSearchConfig{
			Aws:         getBool("ELASTIC_SEARCH_AWS", true),
			Host:        getString("ELASTIC_SEARCH_HOST", ""),
			Sniff:       getBool("ELASTIC_SEARCH_SNIFF", true),
			HealthCheck: getBool("ELASTIC_SEARCH_HEALTH_CHECK", true),
			Debug:       getBool("ELASTIC_SEARCH_DEBUG", false),
			Username:    getString("ELASTIC_SEARCH_USERNAME", ""),
			Password:    getString("ELASTIC_SEARCH_PASSWORD", ""),
		},
	}
}

func getString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func getInt(key string, defaultValue int) int {
	valStr := getString(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}

	return defaultValue
}

func getBool(key string, defaultValue bool) bool {
	valStr := getString(key, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultValue
}

func getDuration(key string, defaultValue int) time.Duration {
	return time.Duration(getInt(key, defaultValue))
}
