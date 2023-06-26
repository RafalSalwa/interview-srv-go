package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type IntrvClient struct {
	connection *Connection
	handlers   map[string]EventHandler
	debug      bool
}

func NewClient(connection *Connection) *IntrvClient {
	return &IntrvClient{
		connection: connection,
		handlers:   make(map[string]EventHandler),
		debug:      false,
	}
}

func (c *IntrvClient) SetDebug(debug bool) {
	c.debug = debug
}

func (c *IntrvClient) SetHandler(eventName string, handler EventHandler) {
	c.handlers[eventName] = handler
}

func (c *IntrvClient) HandleChannel(ctx context.Context, channelName string, consumerName string, args amqp.Table) error {
	consumer, err := c.connection.CreateConsumer(channelName, consumerName, c.handleEvent, args)
	if err != nil {
		return err
	}

	defer consumer.Close()
	return consumer.Handle(ctx)
}

func (c *IntrvClient) handleEvent(data []byte) (isSuccess bool) {
	// create new event and deserialize it
	event := Event{}
	_ = json.Unmarshal(data, &event)

	if c.debug {
		fmt.Printf("processing event %s\n %#v\n", event.Name, event)
	}

	if handler, ok := c.handlers[event.Name]; ok {
		_ = handler(event)
	}
	// we have no handler for that type of event
	return true
}
