package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/pkg/rabbitmq"
)

type AccountPasswordChange struct {
	UserId int `json:"user_id"`
}

func WrapHandleUserPasswordChange(event rabbitmq.Event) error {
	var data AccountPasswordChange
	_ = json.Unmarshal([]byte(event.Content), &data)

	return HandleUserPasswordChange(data)
}

func HandleUserPasswordChange(payload AccountPasswordChange) error {
	fmt.Println("HandleUserPasswordChange", payload)
	return nil
}
