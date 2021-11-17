package config

import (
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/log"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Env                    string
	Port                   int
	Logging                bool
	LogPath                string
	Network                string
	Index                  string
	Debug                  bool
	ElasticSearch          ElasticSearchConfig
	Throttle               ThrottleConfig
	Subscribe              bool
	Cache                  bool
	CacheDefaultExpiration time.Duration
	Aws                    AwsConfig
	SentryDsn              string
}

type AwsConfig struct {
	AccessKey string
	SecretKey string
	Token     string
	Region    string
}

type ElasticSearchConfig struct {
	Aws         bool
	Hosts       []string
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
		zap.L().With(zap.Error(err)).Fatal("Unable to init config")
	}

	initLogger()

	initSentry()
}
func initLogger() {
	log.NewLogger(fmt.Sprintf("%s/indexer.log", Get().LogPath), Get().Debug, Get().SentryDsn)
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
		Env:                    getString("ENV", ""),
		Port:                   getInt("PORT", 8080),
		Logging:                getBool("LOGGING", false),
		LogPath:                getString("LOG_PATH", "/app/logs"),
		Network:                getString("NETWORK", "zilliqa"),
		Index:                  getString("INDEX_NAME", "xxx"),
		Debug:                  getBool("DEBUG", false),
		Subscribe:              getBool("SUBSCRIBE", false),
		Cache:                  getBool("CACHE", false),
		CacheDefaultExpiration: getDuration("CACHE_DEFAULT_EXPIRATION", 60) * time.Second,
		SentryDsn:              getString("SENTRY_DSN", ""),
		Aws: AwsConfig{
			AccessKey: getString("AWS_ACCESS_KEY", ""),
			SecretKey: getString("AWS_SECRET_KEY", ""),
			Token:     getString("AWS_TOKEN", ""),
			Region:    getString("AWS_REGION", ""),
		},
		ElasticSearch: ElasticSearchConfig{
			Aws:         getBool("ELASTIC_SEARCH_AWS", true),
			Hosts:       getSlice("ELASTIC_SEARCH_HOSTS", make([]string, 0), ","),
			Sniff:       getBool("ELASTIC_SEARCH_SNIFF", true),
			HealthCheck: getBool("ELASTIC_SEARCH_HEALTH_CHECK", true),
			Debug:       getBool("ELASTIC_SEARCH_DEBUG", false),
			Username:    getString("ELASTIC_SEARCH_USERNAME", ""),
			Password:    getString("ELASTIC_SEARCH_PASSWORD", ""),
		},
		Throttle: ThrottleConfig{
			MaxEventsPerSec: getInt("MAX_EVENT_PER_SEC", 1000),
			MaxBurstSize:    getInt("MAX_BURST_SIZE", 20),
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

func getSlice(key string, defaultVal []string, sep string) []string {
	valStr := getString(key, "")
	if valStr == "" {
		return defaultVal
	}

	return strings.Split(valStr, sep)
}

func getDuration(key string, defaultValue int) time.Duration {
	return time.Duration(getInt(key, defaultValue))
}
