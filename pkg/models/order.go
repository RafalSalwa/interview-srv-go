package models

import "time"

type OrderDBModel struct {
	Id             int `gorm:"id;primaryKey;autoIncrement"`
	UserId         int `gorm:"column:user_id;type:int;not null;uniqueIndex;not null"`
	SubscriptionId int `gorm:"column:plan_id;type:varchar(255);not null"`
	Price          int
	VoucherId      int
	PurchasedAt    time.Time  `gorm:"column:purchased_at"`
	CreatedAt      time.Time  `gorm:"column:created_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at"`
}

func (OrderDBModel) TableName() string {
	return "orders"
}
