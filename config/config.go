package config

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

type Conf struct {
	App    ConfApp
	Server ConfServer
	DB     ConfDB
	GRPC   ConfGRPC
	AMQP   AMQP
	Token  ConfToken
}

type Mongo struct {
	DBUri string `env:"MONGODB_SERVER_ADDR"`
}

type ConfGRPC struct {
	GrpcServerAddress string `env:"GRPC_SERVER_ADDRESS"`
}

type ConfApp struct {
	Env          string `env:"APP_ENV, default=dev"`
	Debug        bool   `env:"APP_DEBUG, default=false"`
	JwtSecretKey string `env:"JWT_SECRET_KEY, required"`
}

type ConfBasicAuth struct {
	Username string `env:"SERVER_AUTH_BASIC_USERNAME"`
	Password string `env:"SERVER_AUTH_BASIC_PASSWORD"`
}

type ConfServer struct {
	Addr         string        `env:"SERVER_ADDR"`
	Host         string        `env:"SERVER_HOST"`
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	APIKey       string        `env:"X_API_KEY, required"`
	AuthMethod   string        `env:"SERVER_AUTH_METHOD"`
	BasicAuth    ConfBasicAuth
	BearerToken  string `env:"BEARER_TOKEN"`
	Debug        bool   `env:"APP_DEBUG,required"`
}

type AMQP struct {
	Protocol string `env:"AMQP_PROTOCOL"`
	Username string `env:"AMQP_USERNAME"`
	Password string `env:"AMQP_PASSWORD"`
	Hostname string `env:"AMQP_HOSTNAME"`
	VHost    string `env:"AMQP_VHOST"`
}

type ConfDB struct {
	Addr     string `env:"MYSQL_ADDR,required"`
	Username string `env:"MYSQL_USER,required"`
	Password string `env:"MYSQL_PASSWORD,required"`
	DBName   string `env:"MYSQL_NAME,required"`
	Debug    bool   `env:"MYSQL_DEBUG,required"`
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

func readFromFile() (bool, error) {
	re := regexp.MustCompile(`^(.*` + "interview-app-srv" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	err := godotenv.Load(string(rootPath) + `.env`)

	if err != nil {
		fmt.Println("err:", err)
		errL := godotenv.Load(".env")
		if errL != nil {
			fmt.Println("errL:", errL)
		}
	}
	return true, nil
}
