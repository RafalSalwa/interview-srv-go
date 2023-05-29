package amqp

import "context"

type Client interface {
	SetHandler(eventName string, handler EventHandler)
	HandleChannel(ctx context.Context, channelName string, requeueOnError bool) error
}
