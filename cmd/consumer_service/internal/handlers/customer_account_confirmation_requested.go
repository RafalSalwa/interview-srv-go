package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/pkg/encdec"
	"log"
	"net/mail"

	"github.com/RafalSalwa/interview-app-srv/cmd/consumer_service/config"
	"github.com/RafalSalwa/interview-app-srv/pkg/email"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/RafalSalwa/interview-app-srv/pkg/rabbitmq"
)

func WrapHandleCustomerAccountRequestConfirmEmail(event rabbitmq.Event) error {
	var data models.UserEvent
	err := json.Unmarshal([]byte(event.Content), &data)
	if err != nil {
		return err
	}
	c, err := config.InitConfig()
	mailer := email.NewClient(c.Email)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerAccountRequestConfirmEmail(data, mailer)
}

func CustomerAccountRequestConfirmEmail(payload models.UserEvent, mailer email.Client) error {
	addr, err := encdec.Decrypt(payload.Email)
	fmt.Println("addr", addr)
	if err != nil {
		return err
	}
	_, err = mail.ParseAddress(addr)
	if err != nil {
		err = mailer.SendVerificationEmail(email.UserEmailData{
			Username:         payload.Username,
			Email:            addr,
			VerificationCode: payload.VerificationCode,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
