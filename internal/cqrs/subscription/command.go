package subscription

import (
	"time"
)

type Command interface {
	Execute() (interface{}, error)
}

type CreateSubscriptionCommand struct {
	UserId    int        `gorm:"column:user_id;type:int;not null;uniqueIndex;not null"`
	PlanId    int        `gorm:"column:plan_id;type:varchar(255);not null"`
	StartedAt *time.Time `gorm:"column:started_at"`
}

type CreditAccountCommand struct {
	AccountId string `json:"account_id"`
	Amount    int    `json:"amount"`
}

type DebitAccountCommand struct {
	AccountId string `json:"account_id"`
	Amount    int    `json:"amount"`
}

func (command *CreateSubscriptionCommand) Execute() (interface{}, error) {
	// id := MakeAccount(command.OpeningBalance, store).Id
	return 1, nil
}
