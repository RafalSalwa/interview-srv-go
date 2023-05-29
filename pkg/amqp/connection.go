package amqp

import (
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/streadway/amqp"
)

type Connection struct {
	Connection  *amqp.Connection
	credentials Credentials
}

func NewConnectionFromCredentials(credentials config.AMQP) *Connection {
	cred := Credentials{
		Protocol: credentials.Protocol,
		Username: credentials.Username,
		Password: credentials.Password,
		Hostname: credentials.Hostname,
		VHost:    credentials.VHost,
	}

	return &Connection{
		credentials: cred,
	}
}

func (l *Connection) Connect() error {
	c, err := amqp.Dial(l.credentials.GetUrl())
	if err != nil {
		return err
	}
	l.Connection = c

	return nil
}

func (l *Connection) Close() error {
	return l.Connection.Close()
}
