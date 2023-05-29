package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/RafalSalwa/interview-app-srv/pkg/amqp"
)

type AccountPasswordChange struct {
	UserId int `json:"user_id"`
}

func WrapHandleUserPasswordChange(event amqp.Event) error {
	var data AccountPasswordChange
	_ = json.Unmarshal([]byte(event.Content), &data)

	return HandleUserPasswordChange(data)
}

func HandleUserPasswordChange(payload AccountPasswordChange) error {
	fmt.Println("HandleUserPasswordChange", payload)
	return nil
}
