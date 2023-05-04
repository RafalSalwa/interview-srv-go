package models

import "time"

// swagger:model User
type UserDBModel struct {
	Id               int64     `json:"id"`
	Username         string    `json:"username"`
	Password         string    `json:"password,omitempty"`
	Firstname        *string   `json:"firstname"`
	Lastname         *string   `json:"lastname"`
	Email            *string   `json:"email"`
	RolesJson        string    `json:"rolesJson"`
	Roles            []string  `json:"roles"`
	VerificationCode string    `json:"verification_ode,omitempty"`
	Verified         bool      `json:"is_verified"`
	Active           bool      `json:"is_active"`
	JwtToken         *string   `json:"AuthToken"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	LastLogin        time.Time `json:"last_login"`
	DeletedAt        time.Time `json:"deleted_at"`
}

// swagger:model Users
type Users struct {
	Users []UserDBModel `json:"users"`
}

type UserRequest struct {
	Id       int64   `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Role     *string `json:"role"`
}

type CreateUserRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDBResponse struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	Firstname *string   `json:"firstname"`
	Lastname  *string   `json:"lastname"`
	Email     string    `json:"email"`
	RolesJson string    `json:"rolesJson"`
	Roles     []string  `json:"roles"`
	Role      string    `json:"role"`
	Verified  bool      `json:"verified"`
	Active    bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type UserResponse struct {
	Id        int64     `json:"id,omitempty"`
	Username  string    `json:"username"`
	Firstname *string   `json:"firstname"`
	RolesJson string    `json:"rolesJson"`
	Roles     []string  `json:"roles"`
	Verified  bool      `json:"verified"`
	Active    bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	LastLogin time.Time `json:"last_login"`
	DeletedAt time.Time `json:"deleted_at"`
}

type UpdateUserRequest struct {
	Id        int64   `json:"id,omitempty"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Password  *string `json:"password"`
}
