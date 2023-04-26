package models

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
