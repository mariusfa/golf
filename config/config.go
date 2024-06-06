package config

import (
	"errors"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

func GetConfig(filename string, config any) error {
	godotenv.Load(filename) // ignore error. if file does not exist, it's ok

	v := reflect.ValueOf(config).Elem()

	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		fieldName := fieldInfo.Name
		envName := getEnvName(fieldName)
		envValue := os.Getenv(envName)
		requiredTag := fieldInfo.Tag.Get("required")

		if envValue == "" && requiredTag != "false" {
			return errors.New("env var " + envName + " is required")
		}

		v.Field(i).SetString(envValue)
	}
	return nil
}

// Gets env name from field name
// Example: ServerPort -> SERVER_PORT
func getEnvName(fieldName string) string {
	envName := ""
	for i, letter := range fieldName {
		if i > 0 && letter >= 'A' && letter <= 'Z' {
			envName += "_"
		}
		envName += string(letter)
	}
	envName = strings.ToUpper(envName)
	return envName
}
