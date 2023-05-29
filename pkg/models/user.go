package models

import (
	"gorm.io/gorm"
	"time"
)

// swagger:model User
type UserDBModel struct {
	Id               int64      `gorm:"id;primaryKey;autoIncrement"`
	Username         string     `gorm:"type:varchar(180);not null;uniqueIndex;not null"`
	Password         string     `gorm:"type:varchar(255);not null"`
	Firstname        *string    `gorm:"column:first_name;type:varchar(255)"`
	Lastname         *string    `gorm:"column:last_name;type:varchar(255)"`
	Email            string     `gorm:"type:varchar(255);not null;uniqueIndex;not null"`
	Phone            *string    `gorm:"column:phone_no;type:varchar(11)"`
	Roles            []byte     `gorm:"column:roles"`
	VerificationCode string     `gorm:"column:verification_code;type:varchar(6)"`
	Verified         bool       `gorm:"column:is_verified;default:false"`
	Active           bool       `gorm:"column:is_active;default:false"`
	CreatedAt        time.Time  `gorm:"column:created_at"`
	UpdatedAt        *time.Time `gorm:"column:updated_at"`
	LastLogin        *time.Time `gorm:"column:last_login"`
	DeletedAt        *time.Time `gorm:"column:deleted_at"`
}

func (um *UserDBModel) BeforeCreate(tx *gorm.DB) (err error) {

	um.Active = false
	um.Verified = false
	return nil
}

func (UserDBModel) TableName() string {
	return "user"
}

// swagger:model Users
type Users struct {
	Users []UserDBModel `json:"users"`
}

type UserRequest struct {
	Id       int64   `json:"id" govalidator:"int"`
	Username *string `json:"username"`
	Email    *string `json:"email" `
}

type CreateUserRequest struct {
	Username        string `json:"username" validate:"required,min=3,max=32"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,max=16"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8,max=16"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required_without=Email"`
	Email    string `json:"email" validate:"required_without=Username,omitempty,email"`
	Password string `json:"password" validate:"required"`
}

type UserDBResponse struct {
	Id        int64
	Username  string
	Firstname *string
	Lastname  *string
	Email     string
	Password  string
	RolesJson string
	Roles     []string
	Role      string
	Verified  bool
	Active    bool
	CreatedAt time.Time
	LastLogin *time.Time
}

type UserResponse struct {
	Id               int64      `json:"id,omitempty"`
	Username         string     `json:"username"`
	Firstname        *string    `json:"firstname,omitempty"`
	RolesJson        string     `json:"rolesJson,omitempty"`
	Roles            string     `json:"roles,omitempty"`
	Verified         bool       `json:"is_verified,omitempty"`
	VerificationCode string     `json:"verification_token,omitempty"`
	Active           bool       `json:"is_active,omitempty"`
	Token            string     `json:"token,omitempty"`
	RefreshToken     string     `json:"refresh_token,omitempty"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
	LastLogin        *time.Time `json:"last_login,omitempty"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

type UpdateUserRequest struct {
	Id        int     `json:"id,omitempty"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Password  *string `json:"password"`
}

type ChangePasswordRequest struct {
	Id              int    `json:"id"`
	OldPassword     string `json:"oldPassword"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type User struct {
	ID int
	// Creates a primary key for UserID.
	UserID int `gorm:"primary_key"`
	// Creates constrains for Username
	// -> 15 character max limit and not be passed a the blank
	Username string `sql:"type:VARCHAR(15);not null"`
	// Creates constraints for FirstName
	// -> 100 character max limit, Not be passed a the blank, Column name will not be FName, will be FirstName
	FName string `sql:"size:100;not null" gorm:"column:FirstName"`
	// Creates consstraints for LastName
	// -> Unique index/constraint, Not be passed a the blank, Default value is 'Unknown', Column name will not be LName, will be LastName
	LName     string `sql:"unique;unique_index;not null;DEFAULT:'Unknown'" gorm:"column:LastName"`
	Count     int    `gorm:"AUTO_INCREMENT"`
	TempField bool   `sql:"-"` // Ignore a Field
}

type Roles struct {
	Roles []string
}
