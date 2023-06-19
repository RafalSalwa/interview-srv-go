package config

import (
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/pkg/tracing"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/spf13/viper"

	"github.com/joeshaw/envdecode"
	_ "github.com/joho/godotenv"
)

type Conf struct {
	App    ConfApp              `mapstructure:",squash"`
	Server ConfServer           `mapstructure:",squash"`
	DB     ConfDB               `mapstructure:",squash"`
	GRPC   ConfGRPC             `mapstructure:",squash"`
	AMQP   AMQP                 `mapstructure:",squash"`
	Token  ConfToken            `mapstructure:",squash"`
	Jaeger tracing.JaegerConfig `mapstructure:",squash"`
}

type Mongo struct {
	DBUri string `env:"MONGODB_SERVER_ADDR" mapstructure:"MONGODB_SERVER_ADDR"`
}

type ConfGRPC struct {
	GrpcServerAddress string `env:"GRPC_SERVER_ADDRESS" mapstructure:"GRPC_SERVER_ADDRESS"`
}

type ConfApp struct {
	Env           string `env:"APP_ENV, default=dev" mapstructure:"APP_ENV"`
	Debug         bool   `env:"APP_DEBUG, default=false" mapstructure:"APP_DEBUG"`
	JaegerEnabled bool   `env:"JAEGER_ENABLE, default=false" mapstructure:"JAEGER_ENABLE"`
	JwtSecretKey  string `env:"JWT_SECRET_KEY, required" mapstructure:"JWT_SECRET_KEY"`
}

type ConfBasicAuth struct {
	Username string `env:"SERVER_AUTH_BASIC_USERNAME"`
	Password string `env:"SERVER_AUTH_BASIC_PASSWORD"`
}

type ConfServer struct {
	Addr         string        `env:"SERVER_ADDR" mapstructure:"SERVER_ADDR"`
	Host         string        `env:"SERVER_HOST" mapstructure:"SERVER_HOST"`
	Port         int           `env:"SERVER_PORT" mapstructure:"SERVER_PORT"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required" mapstructure:"SERVER_TIMEOUT_READ"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required" mapstructure:"SERVER_TIMEOUT_WRITE"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required" mapstructure:"SERVER_TIMEOUT_IDLE"`
	APIKey       string        `env:"X_API_KEY, required" mapstructure:"X_API_KEY"`
	AuthMethod   string        `env:"SERVER_AUTH_METHOD" mapstructure:"SERVER_AUTH_METHOD"`
	BasicAuth    ConfBasicAuth
	BearerToken  string `env:"BEARER_TOKEN" mapstructure:"BEARER_TOKEN"`
	Debug        bool   `env:"APP_DEBUG,required" mapstructure:"APP_DEBUG"`
}

type AMQP struct {
	Protocol string `env:"AMQP_PROTOCOL"`
	Username string `env:"AMQP_USERNAME"`
	Password string `env:"AMQP_PASSWORD"`
	Hostname string `env:"AMQP_HOSTNAME"`
	VHost    string `env:"AMQP_VHOST"`
}

type ConfDB struct {
	Addr     string `env:"MYSQL_ADDR,required" mapstructure:"MYSQL_ADDR"`
	Username string `env:"MYSQL_USER,required" mapstructure:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD,required" mapstructure:"MYSQL_PASSWORD"`
	DBName   string `env:"MYSQL_NAME,required" mapstructure:"MYSQL_NAME"`
	Debug    bool   `env:"MYSQL_DEBUG,required" mapstructure:"MYSQL_DEBUG"`
}

type ConfToken struct {
	AccessTokenPrivateKey  string        `env:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `env:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `env:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `env:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `env:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `env:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `env:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `env:"REFRESH_TOKEN_MAXAGE"`
}

func New() *Conf {
	var c Conf
	_, err := readFromFile()
	if err != nil {
		log.Fatalf("Failed to read env file: %s", err)
	}
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to read env file via envdecode: %s", err)
	}
	return &c
}

func NewViper() *Conf {
	var c Conf
	viper.SetConfigFile(getEnvPath())
	err := viper.ReadInConfig()
	if err != nil {
		return nil
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("Failed to unmarshal env file: %s", err)
		return nil
	}

	return &c
}
func getEnvPath() string {
	re := regexp.MustCompile(`^(.*` + "interview-app-srv" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	return string(rootPath) + `.env`
}

func readFromFile() (bool, error) {
	err := godotenv.Load(getEnvPath())

	if err != nil {
		fmt.Println("err:", err)
		errL := godotenv.Load(".env")
		if errL != nil {
			fmt.Println("errL:", errL)
		}
	}
	return true, nil
}
