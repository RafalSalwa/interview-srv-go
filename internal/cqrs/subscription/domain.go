package subscription

import (
	"time"

	"github.com/RafalSalwa/interview-app-srv/internal/cqrs"
	"github.com/RafalSalwa/interview-app-srv/pkg/models"
)

func MakeSubscription(balance int, store cqrs.EventStore) models.SubscriptionDBModel {
	// id := uuid.New()
	subscription := &models.SubscriptionDBModel{}
	subscription.UserId = 1
	subscription.PlanId = 1
	subscription.PurchasedAt = time.Now()
	subscription.StartedAt = time.Now()
	subscription.EndsAt = time.Now()

	// account.Update(AccountCreated{id, balance})
	// account.Save()
	return *subscription
}
