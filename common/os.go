package common

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// GetenvOrDefault encapsulate built-in os.Getenv behavior but if key is not present it returns the defaultValue
func GetenvOrDefault(key string, defaultValue string) string {
	str := os.Getenv(key)
	if strings.TrimSpace(str) == "" {
		return defaultValue
	}
	return str
}

// GetenvBoolOrDefault returns the valeu of os.Getenv(key string) value as bool or defaultValue if error
// If the environment variable (key) is not defined, it returns the given defaultValue
// If the environment variable (key) is not a valid bool format, it returns the given defaultValue
// If any error occurring during bool parse, it returns the given defaultValue
func GetenvBoolOrDefault(key string, defaultValue bool) bool {
	str := os.Getenv(key)
	val, err := strconv.ParseBool(str)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetenvIntOrDefault returns the valeu of os.Getenv(key string) value as int or defaultValue if error
// If the environment variable (key) is not defined, it returns the given defaultValue
// If the environment variable (key) is not a valid int format, it returns the given defaultValue
// If any error occurring during int parse, it returns the given defaultValue
func GetenvIntOrDefault(key string, defaultValue int64) int64 {
	str := os.Getenv(key)
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defaultValue
	}
	return val
}

// LocalEnvConfig is used to automatically call the InitLocalEnvConfig method using Dependency Injection
// So, if a func parameter or a struct field depends of LocalEnvConfig, when DI starts, it will call InitLocalEnvConfig as the LocalEnvConfig provider
type LocalEnvConfig struct {
	Initialized bool
}

var localEnvConfig *LocalEnvConfig
var localEnvConfigOnce sync.Once

// InitLocalEnvConfig load a .env file to setup local environment vars
// It's invoked once per application process
func InitLocalEnvConfig() *LocalEnvConfig {
	fmt.Println("InitLocalEnvConfig")
	envName := GetenvOrDefault("ENV_NAME", "local")
	if envName == "local" {
		log.Printf("ENVIRONMENT NAME \u001B[31m(%s)\u001B[0m\n", envName)
		localEnvConfigOnce.Do(func() {
			if err := godotenv.Load(); err == nil {
				log.Println("Env vars loaded from .env file on process", os.Getpid())
				localEnvConfig = &LocalEnvConfig{
					Initialized: true,
				}
			} else {
				log.Println("Skipping .env file")
				localEnvConfig = &LocalEnvConfig{
					Initialized: false,
				}
			}
		})
	}

	return localEnvConfig
}

// SetConfigFromEnvVars builds a struct by setting it fields values using the "var" tag
// Constraints: s interface{} - must be a initialized pointer
// Supported types: String, Boolean, Int, Int8, Int16, Int32 and Int64
func SetConfigFromEnvVars(s interface{}) error {
	v := reflect.ValueOf(s)
	t := v.Type()
	if t.Kind() != reflect.Ptr {
		return errors.New("s must be an pointer")
	}
	e := t.Elem()
	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		if tag, ok := f.Tag.Lookup("env"); ok {
			values := strings.Split(tag, ",")
			if len(values) > 0 {
				fv := v.Elem().FieldByName(f.Name)
				if fv.CanSet() {
					switch k := fv.Kind(); k {
					case reflect.Bool:
						fv.SetBool(GetenvBoolOrDefault(values[0], false))
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						fv.SetInt(GetenvIntOrDefault(values[0], 0))
					default:
						fv.SetString(os.Getenv(values[0]))
					}
				}
			}
		}

	}
	return nil
}
