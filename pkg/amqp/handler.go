package amqp

import (
	"context"
	"encoding/json"
	"fmt"
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

func (c *IntrvClient) HandleChannel(ctx context.Context, channelName string, requeueOnError bool) error {
	consumer, err := c.connection.CreateConsumer(channelName, requeueOnError, c.handleEvent)
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

	fmt.Println("handlers", c.handlers)
	if handler, ok := c.handlers[event.Name]; ok {
		_ = handler(event)
		fmt.Println("handler", event, event.Name)
	}
	// we have no handler for that type of event
	return true
}
