package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"

	ErrorModel = "!!!Error"
	ErrorStyle = "--->"

	// service name
	SvcsUser = "UserService"

	PostgresCtxTimeout = time.Second * 30       // postgres context timeout 30s
	PostgresLogTime    = time.Millisecond * 500 // postgres context timeout 500ms
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string

	BotToken string

	PostgresHost           string
	PostgresPort           int
	PostgresUser           string
	PostgresPassword       string
	PostgresDatabase       string
	PostgresMaxConnections int32

	HTTPPort   string
	HTTPScheme string
}

// Load ...
func Load() Config {
	envFileName := cast.ToString(getOrReturnDefaultValue("ENV_FILE_PATH", "./.env"))

	if err := godotenv.Load(envFileName); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "anonymous_letter_postgres"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0.0"))

	config.BotToken = cast.ToString(getOrReturnDefaultValue("BOT_TOKEN", "123123abc"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "0.0.0.0"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "postgres"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", config.ServiceName))

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", "3000"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 80))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
