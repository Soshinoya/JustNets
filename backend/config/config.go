package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}

func GetDSN() (dsn string) {
	dsn = fmt.Sprintf("host=%s user=%s password=%s port=%d sslmode=%s TimeZone=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "root"),
		getEnvAsInt("DB_PORT", 5432),
		getEnv("DB_SSL", "disable"),
		getEnv("DB_TIMEZONE", "EETDST"))
	return dsn + " dbname=%s"
}

func GetLogPath() string {
	return getEnv("LOG_PATH", "C:\\Users\\Shirt\\Desktop\\JustNets\\logs") + "\\%s\\%s-%s.log"
}
