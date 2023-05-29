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
	Hostname string
	VHost    string
}

func NewCredentials() *Credentials {
	return &Credentials{}
}
func NewCredentialsFromConfig() (*Credentials, error) {
	c := &Credentials{
		Protocol: os.Getenv("AMQP_PROTOCOL"),
		Username: os.Getenv("AMQP_USERNAME"),
		Password: os.Getenv("AMQP_PASSWORD"),
		Hostname: os.Getenv("AMQP_HOSTNAME"),
		VHost:    os.Getenv("AMQP_VHOST"),
	}

	if c.Protocol == "" {
		c.Protocol = "amqp"
	}

	if c.Hostname == "" {
		return nil, CredentialsFromEnvHostnameNotSetError
	}

	if c.Hostname == "" {
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
