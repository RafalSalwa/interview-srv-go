package amqp

import (
	"github.com/streadway/amqp"
)

type Connection struct {
	Connection  *amqp.Connection
	credentials Credentials
}

func NewConnection(cfg Config) *Connection {
	cred := Credentials{
		Username: cfg.Username,
		Password: cfg.Password,
		Addr:     cfg.Addr,
        VHost: 
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
