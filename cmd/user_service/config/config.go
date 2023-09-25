package config

import (
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/pkg/env"
	"os"
	"strings"

	"github.com/RafalSalwa/interview-app-srv/pkg/email"
	"github.com/RafalSalwa/interview-app-srv/pkg/rabbitmq"

	"github.com/RafalSalwa/interview-app-srv/pkg/grpc"
	"github.com/RafalSalwa/interview-app-srv/pkg/jwt"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	mongodb "github.com/RafalSalwa/interview-app-srv/pkg/mongo"
	"github.com/RafalSalwa/interview-app-srv/pkg/probes"
	"github.com/RafalSalwa/interview-app-srv/pkg/redis"
	"github.com/RafalSalwa/interview-app-srv/pkg/sql"
	"github.com/RafalSalwa/interview-app-srv/pkg/tracing"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string               `mapstructure:"serviceName"`
	App         App                  `mapstructure:"app"`
	Logger      *logger.Config       `mapstructure:"logger"`
	GRPC        grpc.Config          `mapstructure:"grpc"`
	JWTToken    jwt.JWTConfig        `mapstructure:"jwt"`
	MySQL       sql.MySQL            `mapstructure:"mysql"`
	Mongo       mongodb.Config       `mapstructure:"mongo"`
	Redis       *redis.Config        `mapstructure:"redis"`
	Rabbit      rabbitmq.Config      `mapstructure:"rabbitmq"`
	Probes      probes.Config        `mapstructure:"probes"`
	Jaeger      tracing.JaegerConfig `mapstructure:"jaeger"`
	Mail        email.Config         `mapstructure:"email"`
}

type App struct {
	Env   string `mapstructure:"env"`
	Debug bool   `mapstructure:"debug"`
}

func InitConfig() (*Config, error) {
	cfg := &Config{}
	path, err := env.GetPath("user_service")
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
	appEnv := os.Getenv("APP_ENV")
	if err != nil {
		return "", errors.Wrap(err, "os.Getwd")
	}

	configPath := ""
	if strings.Contains(getwd, "user_service") {
		configPath = fmt.Sprintf("%s/config.%s.yaml", getwd, appEnv)
	} else {
		configPath = fmt.Sprintf("%s/cmd/user_service/config/config.%s.yaml", getwd, appEnv)
	}

	return configPath, nil
}
