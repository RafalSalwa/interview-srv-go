package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Connection struct {
	Connection  *amqp.Connection
	Channel     *amqp.Channel
	credentials Credentials
}

func NewConnection(cfg Config) *Connection {
	cred := Credentials{
		Username: cfg.Username,
		Password: cfg.Password,
		Addr:     cfg.Addr,
		VHost:    cfg.VHost,
		Exchange: cfg.Exchange,
	}

	return &Connection{
		credentials: cred,
	}
}

func (l *Connection) Connect() error {
	c, err := amqp.Dial(l.credentials.GetURL())
	if err != nil {
		return err
	}
	ch, err := c.Channel()
	if err != nil {
		return err
	}
	l.Connection = c
	l.Channel = ch
	// if l.credentials.Exchange != nil {
	//	if err := ch.ExchangeDeclare(
	//		l.credentials.Exchange.Name,
	//		l.credentials.Exchange.Type,
	//		l.credentials.Exchange.Durable,
	//		false,
	//		false,
	//		false,
	//		nil,
	//	); err != nil {
	//		return err
	//	}
	//	if l.credentials.Exchange.Queue != "" {
	//		if _, err := ch.QueueDeclare(
	//			l.credentials.Exchange.Queue,
	//			true,
	//			false,
	//			false,
	//			false,
	//			nil,
	//		); err != nil {
	//			return err
	//		}
	//	}
	//}
	return nil
}

func (l *Connection) Close() error {
	return l.Connection.Close()
}
