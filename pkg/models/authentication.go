package models

import (
	"time"

	"github.com/RafalSalwa/interview-app-srv/util/jwt"
)

type Authentication struct {
	User  *User     `json:"user"`
	Token jwt.Token `json:"token"`
}

type UserDevice struct {
	Id            int64      `json:"id"`
	UserId        int64      `json:"userId"`
	FirebaseToken string     `json:"firebaseToken"`
	OsType        int64      `json:"osType"`
	SdkVersion    int64      `json:"sdkVersion"`
	Model         string     `json:"model"`
	Brand         string     `json:"brand"`
	CreatedAt     time.Time  `json:"createdAt"`
	LastLogin     time.Time  `json:"lastLogin"`
	DeletedAt     *time.Time `json:"deletedAt"`
}
type UserDevices struct {
	Items []UserDevice `json:"items"`
}
