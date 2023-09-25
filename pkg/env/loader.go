package env

import (
	"fmt"
	"os"
	"strings"
)

func GetPath(suffix string) (string, error) {
	getwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get workdir failed: %w", err)
	}

	appEnv := getEnv("APP_ENV", "dev")
	configFile := ""
	if strings.HasSuffix(getwd, suffix) {
		configFile = fmt.Sprintf("%s/config.%s.yaml", getwd, appEnv)
	} else {
		splitted := strings.Split(getwd, suffix)
		configFile = fmt.Sprintf("%s/cmd/%s/config/config.%s.yaml", splitted[0], suffix, appEnv)
	}
	return configFile, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
