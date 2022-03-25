package config

import (
	"github.com/ZilDuck/indexer-api/internal/log"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

var config *Config

type Config struct {
	Env           string
	Port          int
	Debug         bool
	ElasticSearch ElasticSearchConfig
	DBConfig      DBConfig
	Aws           AwsConfig
	CdnHost       string
	AuditDir      string
	ApiClients    []string
}

type AwsConfig struct {
	AccessKey string
	SecretKey string
	Token     string
	Region    string
}

type ElasticSearchConfig struct {
	Host        string
	Sniff       bool
	HealthCheck bool
	Debug       bool
	Username    string
	Password    string
}

type DBConfig struct {
	ConnString string
	LogMode    bool
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	log.NewLogger(Get().Debug)
}

func Get() *Config {
	if config == nil {
		config = &Config{
			Env:   getString("ENV", ""),
			Port:  getInt("PORT", 8082),
			Debug: getBool("DEBUG", false),
			Aws: AwsConfig{
				AccessKey: getString("AWS_ACCESS_KEY_ID", ""),
				SecretKey: getString("AWS_SECRET_KEY_ID", ""),
				Token:     getString("AWS_TOKEN", ""),
				Region:    getString("AWS_REGION", ""),
			},
			ElasticSearch: ElasticSearchConfig{
				Host:        getString("ELASTIC_SEARCH_HOST", ""),
				Sniff:       getBool("ELASTIC_SEARCH_SNIFF", false),
				HealthCheck: getBool("ELASTIC_SEARCH_HEALTH_CHECK", false),
				Debug:       getBool("ELASTIC_SEARCH_DEBUG", false),
				Username:    getString("ELASTIC_SEARCH_USERNAME", ""),
				Password:    getString("ELASTIC_SEARCH_PASSWORD", ""),
			},
			DBConfig: DBConfig{
				ConnString: getString("DB_CONN_STRING", ""),
				LogMode:    getBool("DB_LOG_MODE", false),
			},
			CdnHost:    getString("CDN_HOST", ""),
			AuditDir:   getString("AUDIT_DIR", "/app/audit"),
			ApiClients: getSlice("API_ClIENTS", []string{}, "^"),
		}
	}
	return config
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