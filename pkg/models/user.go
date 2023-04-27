package models

import "time"

// swagger:model User
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

// swagger:model Users
type Users struct {
	Items []User `json:"items"`
}

type SignUpInput struct {
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"passwordConfirm"`
	Role            string    `json:"role"`
	Verified        bool      `json:"verified"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDBResponse struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"passwordConfirm,omitempty"`
	Role            string    `json:"role"`
	Verified        bool      `json:"verified"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type UserResponse struct {
	Id        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
