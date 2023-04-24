package models

import (
	"time"

	"github.com/RafalSalwa/interview-app-srv/api/resource/jwt"
)

type Authentication struct {
	User  *User     `json:"user"`
	Token jwt.Token `json:"token"`
}

type User struct {
	Id            int64        `json:"id"`
	Username      string       `json:"username"`
	Password      string       `json:"password,omitempty"`
	Firstname     *string      `json:"firstname"`
	Lastname      *string      `json:"lastname"`
	RolesJson     string       `json:"rolesJson"`
	Roles         []string     `json:"roles"`
	FirebaseToken *string      `json:"firebaseToken"`
	JwtToken      *string      `json:"AuthToken"`
	Devices       []UserDevice `json:"devices,omitempty"`
}

type Users struct {
	Items []User `json:"items"`
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
