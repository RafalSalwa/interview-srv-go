package models

import "time"

type SubscriptionDBModel struct {
	Id          int       `gorm:"id;primaryKey;autoIncrement"`
	UserId      int       `gorm:"column:user_id;type:int;not null;uniqueIndex;not null"`
	PlanId      int       `gorm:"column:plan_id;type:varchar(255);not null"`
	PurchasedAt time.Time `gorm:"column:purchased_at"`
	StartedAt   time.Time `gorm:"column:started_at"`
	EndsAt      time.Time `gorm:"column:ends_at"`
}

func (SubscriptionDBModel) TableName() string {
	return "subscription"
}
