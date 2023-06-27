package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/pkg/rabbitmq"
)

type CustomerAccountActivatedEventEmail struct {
	UserId         int `json:"user_id"`
	SubscriptionId int `json:"subscription_id"`
}

func WrapHandleUserSubscription(event rabbitmq.Event) error {
	fmt.Println("WrapHandleUserSubscription", event.Name, event.Channel)
	var data CustomerAccountActivatedEventEmail
	_ = json.Unmarshal([]byte(event.Content), &data)

	return HandleUserSubscription(data)
}

func HandleUserSubscription(payload CustomerAccountActivatedEventEmail) error {
	fmt.Println("HandleUserSubscription", payload)
	return nil
}
