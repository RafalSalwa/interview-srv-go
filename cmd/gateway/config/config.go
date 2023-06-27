package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/RafalSalwa/interview-app-srv/pkg/auth"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/pkg/probes"
	"github.com/RafalSalwa/interview-app-srv/pkg/tracing"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string                `mapstructure:"serviceName"`
	App         App                   `mapstructure:"app"`
	Logger      *logger.Config        `mapstructure:"logger"`
	Http        Http                  `mapstructure:"http"`
	Auth        auth.Auth             `mapstructure:"auth"`
	Grpc        Grpc                  `mapstructure:"grpc"`
	Probes      probes.Config         `mapstructure:"probes"`
	Jaeger      *tracing.JaegerConfig `mapstructure:"jaeger"`
}

type App struct {
	Env   string `mapstructure:"env"`
	Debug bool   `mapstructure:"debug"`
}

type Http struct {
	Addr                string        `mapstructure:"addr"`
	Development         bool          `mapstructure:"development"`
	BasePath            string        `mapstructure:"basePath"`
	DebugHeaders        bool          `mapstructure:"debugHeaders"`
	HttpClientDebug     bool          `mapstructure:"httpClientDebug"`
	DebugErrorsResponse bool          `mapstructure:"debugErrorsResponse"`
	IgnoreLogUrls       []string      `mapstructure:"ignoreLogUrls"`
	TimeoutRead         time.Duration `mapstructure:"SERVER_TIMEOUT_READ"`
	TimeoutWrite        time.Duration `mapstructure:"SERVER_TIMEOUT_WRITE"`
	TimeoutIdle         time.Duration `mapstructure:"SERVER_TIMEOUT_IDLE"`
}

type Grpc struct {
	AuthServicePort string `mapstructure:"authServicePort"`
	UserServicePort string `mapstructure:"userServicePort"`
}

func InitConfig() (*Config, error) {
	cfg := &Config{}
	path, err := getEnvPath()
	if err != nil {
		return nil, err
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}
	return cfg, nil
}

func getEnvPath() (string, error) {
	getwd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "os.Getwd")
	}
	configPath := ""
	if strings.Contains(getwd, "gateway") {
		configPath = fmt.Sprintf("%s/config.yaml", getwd)
	} else {
		configPath = fmt.Sprintf("%s/cmd/gateway/config/config.yaml", getwd)
	}
	return configPath, nil
}
