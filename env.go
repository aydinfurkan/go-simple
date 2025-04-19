package simple

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var envLoadedBefore = false

func LazyLoadEnv() {
	if envLoadedBefore {
		return
	}

	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}
	envLoadedBefore = true
}

type Env struct {
	Key   string
	Value string
}

func GetEnv(key string) *Env {
	LazyLoadEnv()

	if value, exists := os.LookupEnv(key); exists {
		return &Env{
			Key:   key,
			Value: value,
		}
	}
	return nil
}

func GetEnvOrThrow(key string) *Env {
	if env := GetEnv(key); env != nil {
		return env
	}
	panic("Environment variable " + key + " not found")
}

func (e *Env) AsString(defaultValue string) string {
	if e == nil {
		return defaultValue
	}
	if e.Value != "" {
		return e.Value
	}
	return defaultValue
}

func (e *Env) AsInt(defaultValue int) int {
	if e == nil {
		return defaultValue
	}

	if intValue, err := strconv.Atoi(e.Value); err == nil {
		return intValue
	}
	return defaultValue
}

func (e *Env) AsInt64(defaultValue int64) int64 {
	if e == nil {
		return defaultValue
	}

	if intValue, err := strconv.ParseInt(e.Value, 10, 64); err == nil {
		return intValue
	}
	return defaultValue
}

func (e *Env) AsFloat64(defaultValue float64) float64 {
	if e == nil {
		return defaultValue
	}

	if floatValue, err := strconv.ParseFloat(e.Value, 64); err == nil {
		return floatValue
	}
	return defaultValue
}

func (e *Env) AsBool(defaultValue bool) bool {
	if e == nil {
		return defaultValue
	}

	if boolValue, err := strconv.ParseBool(e.Value); err == nil {
		return boolValue
	}
	return defaultValue
}

func (e *Env) AsBytes(defaultValue string) []byte {
	if e == nil {
		return []byte(defaultValue)
	}
	if e.Value != "" {
		return []byte(e.Value)
	}
	return []byte(defaultValue)
}

func (e *Env) GetEnvAsStrings(defaultValue string, seperator string) []string {

	var v string

	if e == nil {
		v = defaultValue
	} else {
		v = e.Value
	}

	slc := strings.Split(v, seperator)
	for i := range slc {
		slc[i] = strings.TrimSpace(slc[i])
	}
	return slc
}
