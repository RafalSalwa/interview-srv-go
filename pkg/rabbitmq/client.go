package rabbitmq

import (
	"context"
	"github.com/streadway/amqp"
)

type Client interface {
	SetHandler(eventName string, handler EventHandler)
	HandleChannel(ctx context.Context, channelName string, consumerName string, args amqp.Table) error
}
