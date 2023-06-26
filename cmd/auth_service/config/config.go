package config

import (
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/pkg/email"
	"github.com/RafalSalwa/interview-app-srv/pkg/rabbitmq"
	"os"

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

var configPath string

type Config struct {
	ServiceName string                `mapstructure:"serviceName"`
	App         App                   `mapstructure:"app"`
	Logger      *logger.Config        `mapstructure:"logger"`
	GRPC        grpc.Config           `mapstructure:"grpc"`
	Email       email.Config          `mapstructure:"email"`
	JWTToken    jwt.JWTConfig         `mapstructure:"jwt"`
	MySQL       sql.MySQL             `mapstructure:"mysql"`
	Mongo       mongodb.Config        `mapstructure:"mongo"`
	Redis       *redis.Config         `mapstructure:"redis"`
	Rabbit      rabbitmq.Config       `mapstructure:"rabbitmq"`
	Probes      probes.Config         `mapstructure:"probes"`
	Jaeger      *tracing.JaegerConfig `mapstructure:"jaeger"`
}

type App struct {
	Env   string `mapstructure:"env"`
	Debug bool   `mapstructure:"debug"`
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
	configPath = fmt.Sprintf("%s/cmd/auth_service/config/config.yaml", getwd)

	return configPath, nil
}
