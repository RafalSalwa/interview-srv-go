package rabbitmq

type Event struct {
	Name       string `json:"name"`
	Id         string `json:"id"`
	SequenceId string `json:"seq"`
	TimeStamp  string `json:"ts"`
	Content    string `json:"content"`
	Persist    string `json:"store"`
	Channel    string
}

type EventHandler func(event Event) error
