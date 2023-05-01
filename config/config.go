package config

import (
	"log"
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
}

type ConfGRPC struct {
	GrpcServerAddress string `env:"GRPC_SERVER_ADDRESS"`
}

type ConfApp struct {
	Env          string `env:"APP_ENV, default=dev"`
	Debug        bool   `env:"APP_DEBUG, default=false"`
	JwtSecretKey string `env:"JWT_SECRET_KEY, required"`
}

type ConfServer struct {
	Host         string        `env:"SERVER_HOST,required"`
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	Debug        bool          `env:"APP_DEBUG,required"`
}

type ConfDB struct {
	Host     string `env:"MYSQL_HOSTS,required"`
	Port     int    `env:"MYSQL_PORT,required"`
	Username string `env:"MYSQL_USER,required"`
	Password string `env:"MYSQL_PASSWORD,required"`
	DBName   string `env:"MYSQL_NAME,required"`
	Debug    bool   `env:"MYSQL_DEBUG,required"`
}

func New() *Conf {
	var c Conf
	_, err := readFromFile()
	if err != nil {
		log.Fatalf("Failed to read env file: %s", err)
	}

	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to read env file: %s", err)
	}

	return &c
}

func readFromFile() (bool, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return false, err
	}
	return true, nil
}
