package amqp

import (
	"fmt"
	"os"

	"github.com/RafalSalwa/interview-app-srv/config"
)

var CredentialsFromEnvHostnameNotSetError = fmt.Errorf("cannot create credentials from env. Key AMQP_HOSTNAME is missing")
var CredentialsFromEnvVHostNotSetError = fmt.Errorf("cannot create credentials from env. Key AMQP_VHOST is missing")

type Credentials struct {
	Protocol string
	Username string
	Password string
	Addr     string
	VHost    string
}

func NewCredentialsFromConfig() (*Credentials, error) {
	c := &Credentials{
		Username: os.Getenv("AMQP_USERNAME"),
		Password: os.Getenv("AMQP_PASSWORD"),
		Addr:     os.Getenv("AMQP_HOSTNAME"),
	}

	c.Protocol = "amqp"

	if c.Hostname == "" {
		return nil, CredentialsFromEnvHostnameNotSetError
	}

	if c.VHost == "" {
		return nil, CredentialsFromEnvVHostNotSetError
	}

	return c, nil
}

func (c *Credentials) GetUrl() string {
	result := c.Protocol + "://"
	if c.Username != "" && c.Password != "" {
		result = result + c.Username + ":" + c.Password + "@"
	}
	result = result + c.Hostname + "/" + c.VHost
	return result
}

func (c *Credentials) FromConfig(conf config.AMQP) {
	c.Protocol = conf.Protocol
	c.Username = conf.Username
	c.Password = conf.Password
	c.Hostname = conf.Hostname
	c.VHost = conf.VHost
}
