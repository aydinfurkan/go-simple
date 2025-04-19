package simple

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetEnvOrThrow(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic("Environment variable " + key + " not found")
}

func GetEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func GetEnvAsBytes(key, defaultValue string) []byte {
	if value, exists := os.LookupEnv(key); exists {
		return []byte(value)
	}
	return []byte(defaultValue)
}

func GetEnvAsStrings(key, defaultValue string, seperator string) []string {
	v := defaultValue
	if value, exists := os.LookupEnv(key); exists {
		v = value
	}

	slc := strings.Split(v, seperator)
	for i := range slc {
		slc[i] = strings.TrimSpace(slc[i])
	}
	return slc
}
