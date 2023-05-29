package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/RafalSalwa/interview-app-srv/pkg/amqp"
)

type CustomerAccountDeletion struct {
	UserId int `json:"user_id"`
}

func WrapHandleUserAccountDeletion(event amqp.Event) error {
	var data CustomerAccountDeletion
	_ = json.Unmarshal([]byte(event.Content), &data)

	return HandleUserAccountDeletion(data)
}

func HandleUserAccountDeletion(payload CustomerAccountDeletion) error {
	fmt.Println("HandleUserAccountDeletion", payload)
	return nil
}
