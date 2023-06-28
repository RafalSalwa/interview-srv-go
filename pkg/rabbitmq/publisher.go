package rabbitmq

import (
    "context"
    "github.com/streadway/amqp"
)

type Publisher struct {
    Connection *Connection
    Exchange   *Exchange
}

func NewPublisher(cfg Config) (*Publisher, error) {
    con := NewConnection(cfg)
    if err := con.Connect(); err != nil {
        return nil, err
    }
    return &Publisher{Connection: con, Exchange: cfg.Exchange}, nil
}

func (p *Publisher) Disconnect() error {
    return p.Connection.Close()
}

func (p *Publisher) Publish(ctx context.Context, mes amqp.Publishing) error {
    
    err := p.Connection.Channel.Publish(
        p.Exchange.Name,
        p.Exchange.RoutingKey,
        false,
        false,
        mes,
    )
    if err != nil {
        return err
    }

    return nil
}
