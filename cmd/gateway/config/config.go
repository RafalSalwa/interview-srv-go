package config

import (
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/pkg/csrf"
	"github.com/RafalSalwa/interview-app-srv/pkg/env"
	"github.com/RafalSalwa/interview-app-srv/pkg/http"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"os"
	"strings"

	"github.com/RafalSalwa/interview-app-srv/pkg/http/auth"
	"github.com/RafalSalwa/interview-app-srv/pkg/probes"
	"github.com/RafalSalwa/interview-app-srv/pkg/tracing"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string               `mapstructure:"serviceName"`
	App         App                  `mapstructure:"app"`
	Logger      *logger.Config       `mapstructure:"logger"`
	Http        http.Config          `mapstructure:"http"`
	Auth        auth.Auth            `mapstructure:"auth"`
	Grpc        Grpc                 `mapstructure:"grpc"`
	Probes      probes.Config        `mapstructure:"probes"`
	Jaeger      tracing.JaegerConfig `mapstructure:"jaeger"`
	CSRF        csrf.Config          `mapstructure:"csrf"`
}

type App struct {
	Env   string `mapstructure:"env"`
	Debug bool   `mapstructure:"debug"`
}

type Grpc struct {
	AuthServicePort string `mapstructure:"authServicePort"`
	UserServicePort string `mapstructure:"userServicePort"`
}

func InitConfig() (*Config, error) {
	path, err := env.GetPath("gateway")
	if err != nil {
		return nil, err
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err = viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func getEnvPath() (string, error) {
	getwd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "os.Getwd")
	}
	appEnv := getEnv("APP_ENV", "dev")

	configPath := ""
	if strings.HasSuffix(getwd, "gateway") {
		configPath = fmt.Sprintf("%s/config.%s.yaml", getwd, appEnv)
	} else {
		splitted := strings.Split(getwd, "gateway")
		configPath = fmt.Sprintf("%s/cmd/gateway/config/config.%s.yaml", splitted[0], appEnv)
	}
	return configPath, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
