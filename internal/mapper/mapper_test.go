package mapper

import (
	"testing"
	"time"

	"github.com/RafalSalwa/interview-app-srv/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestMapUserDBModelToUserResponse(t *testing.T) {
	str := "firstname"
	now := time.Now()
	expected := &models.UserResponse{
		Id:           1,
		Username:     "username",
		Firstname:    &str,
		RolesJson:    "{\"roles\": [\"ROLE_ADMIN\", \"ROLE_USER\"]}",
		Roles:        []string{"ROLE_ADMIN", "ROLE_USER"},
		Verified:     true,
		Active:       true,
		Token:        "",
		RefreshToken: "",
		CreatedAt:    &now,
		LastLogin:    &now,
		UpdatedAt:    &time.Time{},
		DeletedAt:    &now,
	}
	dbu := &models.UserDBModel{
		Id:        1,
		Username:  "username",
		Firstname: &str,
		Lastname:  &str,
		Email:     "",
		Password:  "",
		RolesJson: "{\"roles\": [\"ROLE_ADMIN\", \"ROLE_USER\"]}",
		Verified:  true,
		Active:    true,
		CreatedAt: now,
		LastLogin: now,
		DeletedAt: now,
	}
	ur := MapUserDBModelToUserResponse(dbu)
	assert.Equal(t, expected, ur)
}

func TestMapUserDBResponseToUserResponse(t *testing.T) {
	str := "firstname"
	now := time.Now()
	dbu := &models.UserDBResponse{
		Id:        1,
		Username:  "username",
		Firstname: &str,
		Lastname:  &str,
		Email:     "",
		Password:  "",
		RolesJson: "{\"roles\": [\"ROLE_ADMIN\", \"ROLE_USER\"]}",
		Verified:  true,
		Active:    true,
		CreatedAt: now,
		LastLogin: &now,
	}

	expected := &models.UserResponse{
		Id:           1,
		Username:     "username",
		Firstname:    &str,
		RolesJson:    "{\"roles\": [\"ROLE_ADMIN\", \"ROLE_USER\"]}",
		Roles:        []string{"ROLE_ADMIN", "ROLE_USER"},
		Verified:     true,
		Active:       true,
		Token:        "",
		RefreshToken: "",
		CreatedAt:    &now,
		LastLogin:    &now,
	}

	ur := MapUserDBResponseToUserResponse(dbu)
	assert.Equal(t, expected, ur)
}
