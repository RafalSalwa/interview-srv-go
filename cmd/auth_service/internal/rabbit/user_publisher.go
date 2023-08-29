package rabbit

import (
	"context"
	"github.com/RafalSalwa/interview-app-srv/pkg/rabbitmq"
	"github.com/streadway/amqp"
	"go.opentelemetry.io/otel"
	otelcodes "go.opentelemetry.io/otel/codes"
)

type Publisher struct {
	Connection *rabbitmq.Connection
	Exchange   *rabbitmq.Exchange
}

func NewPublisher(cfg rabbitmq.Config) (*Publisher, error) {
	con := rabbitmq.NewConnection(cfg)
	if err := con.Connect(); err != nil {
		return nil, err
	}
	return &Publisher{Connection: con, Exchange: cfg.Exchange}, nil
}

func (p *Publisher) Disconnect() error {
	return p.Connection.Close()
}

func (p *Publisher) Publish(ctx context.Context, mes amqp.Publishing) error {
	ctx, span := otel.GetTracerProvider().Tracer("auth_service-rabbit").Start(ctx, "AMQP Publish user SignUp")
	defer span.End()
	err := p.Connection.Channel.Publish(
		p.Exchange.Name,
		p.Exchange.RoutingKey,
		false,
		false,
		mes,
	)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelcodes.Error, err.Error())
		return err
	}

	return nil
}
