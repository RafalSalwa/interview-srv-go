package event

import (
	"encoding/json"

	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Prepare(ue models.UserEvent) amqp.Publishing {
	body, _ := json.Marshal(ue)
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}
	return message
}
