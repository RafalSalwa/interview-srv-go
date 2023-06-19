package cqrs

import "reflect"

type Aggregate struct {
	Id       string
	name     string
	Version  int
	Events   []Event
	entity   interface{}
	handlers Handlers
	store    EventStore
}

type EventMetaData struct {
	Id               string `json:"id"`
	OccuredAt        string `json:"occurred_at"`
	AggregateVersion int    `json:"aggregate_version"`
	AggregateName    string `json:"aggregate_name"`
	Type             string `json:"type"`
}

type Event struct {
	EventMetaData
	Payload interface{} `json:"payload"`
}

type Handlers map[reflect.Type]func(interface{}, interface{})

type EventStore interface {
	GetEvent(id string) Event
	GetEvents(aggregateId string) []Event
	SaveEvents(aggregateId string, events []Event) error
	GetEventMetaDataFrom(offset, count int) []EventMetaData
}
