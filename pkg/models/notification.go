package models

import (
	"time"
)

type NotificationData struct {
	Sensor int64 `json:"sensor_id"`
	UserId int64 `json: "user_id`
}
type Notification struct {
	Id          int64      `json:"id"`
	UserId      *int64     `json:"userId"`
	Title       string     `json:"title"`
	Body        *string    `json:"body"`
	ClickAction *string    `json:"clickAction"`
	ChannelId   *int64     `json:"channelId"`
	Data        *string    `json:"data" db:"data"`
	DeviceId    int64      `json:"device_id"`
	Sent        int64      `json:"sent"`
	CreatedAt   time.Time  `json:"createdAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}

type Notifications struct {
	Items []*Notification `json:"items"`
}
