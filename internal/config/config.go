package config

import (
	"fmt"
	"github.com/dantudor/zilkroad-txapi/internal/log"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port          int
	Logging       bool
	LogPath       string
	Network       string
	Index         string
	Debug         bool
	ElasticSearch ElasticSearchConfig
}

type ElasticSearchConfig struct {
	Hosts       []string
	Sniff       bool
	HealthCheck bool
	Debug       bool
	Username    string
	Password    string
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		zap.L().With(zap.Error(err)).Fatal("Unable to init config")
	}

	initLogger()
}

func initLogger() {
	log.NewLogger(fmt.Sprintf("%s/api.log", Get().LogPath), Get().Debug)
}

func Get() *Config {
	return &Config{
		Port:    getInt("PORT", 8080),
		Logging: getBool("LOGGING", false),
		LogPath: getString("LOG_PATH", "/app/logs"),
		Network: getString("NETWORK", "zilliqa"),
		Index:   getString("INDEX_NAME", "xxx"),
		Debug:   getBool("DEBUG", false),
		ElasticSearch: ElasticSearchConfig{
			Hosts:       getSlice("ELASTIC_SEARCH_HOSTS", make([]string, 0), ","),
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

func getSlice(key string, defaultVal []string, sep string) []string {
	valStr := getString(key, "")
	if valStr == "" {
		return defaultVal
	}

	return strings.Split(valStr, sep)
}
