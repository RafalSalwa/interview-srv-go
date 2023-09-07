package event

import (
	"encoding/json"

	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/streadway/amqp"
)

func Prepare(ue models.UserEvent) amqp.Publishing {
	body, _ := json.Marshal(ue)
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}
	return message
}
