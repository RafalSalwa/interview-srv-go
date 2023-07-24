package models

import (
	"github.com/RafalSalwa/interview-app-srv/pkg/jwt"
	intrvproto "github.com/RafalSalwa/interview-app-srv/proto/grpc"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type usersModels struct {
	uresp   UserResponse
	ureq    UserRequest
	udbResp UserDBResponse
	udb     UserDBModel
}

func initStructs() usersModels {
	uid := int64(1)
	username := "username"
	email := username + "@interview.com"
	firstname := "firstname"
	lastname := "lastname"
	password := "password"
	phoneNo := "123123123"
	VerificationCode := "abcdefgh"
	timeNow := time.Now()
	uresp := UserResponse{
		Id:               uid,
		Username:         username,
		Firstname:        firstname,
		Verified:         true,
		VerificationCode: "",
		Active:           true,
		Token:            "",
		RefreshToken:     "",
		CreatedAt:        timeNow,
		UpdatedAt:        timeNow,
		LastLogin:        timeNow,
		DeletedAt:        timeNow,
	}
	ureq := UserRequest{
		Id:       uid,
		Username: username,
		Email:    email,
	}
	udbresp := UserDBResponse{
		Id:        uid,
		Username:  username,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Password:  password,
		Verified:  true,
		Active:    true,
		CreatedAt: timeNow,
		LastLogin: timeNow,
	}
	udb := UserDBModel{
		Id:               uid,
		Username:         username,
		Password:         password,
		Firstname:        firstname,
		Lastname:         lastname,
		Email:            email,
		Phone:            phoneNo,
		VerificationCode: VerificationCode,
		Verified:         true,
		Active:           true,
		CreatedAt:        timeNow,
		UpdatedAt:        &timeNow,
		LastLogin:        &timeNow,
		DeletedAt:        &timeNow,
	}
	um := usersModels{
		uresp:   uresp,
		ureq:    ureq,
		udbResp: udbresp,
		udb:     udb,
	}
	return um
}

func TestUserResponseMappers(t *testing.T) {
	um := initStructs()
	ur := UserResponse{}
	err := ur.FromDBModel(&um.udb)
	assert.NoError(t, err)
	ur = UserResponse{}
	err = ur.FromDBResponse(&um.udbResp)
	assert.NoError(t, err)
	tp := &jwt.TokenPair{
		AccessToken:  "access",
		RefreshToken: "refresh",
	}
	ur.AssignTokenPair(tp)
	assert.NoError(t, err)
	assert.Equal(t, ur.Token, "access")
	assert.Equal(t, ur.RefreshToken, "refresh")

	ur = UserResponse{}
	ur.FromProtoSignIn(&intrvproto.SignInUserResponse{
		Status:       "ok",
		AccessToken:  "access",
		RefreshToken: "refresh",
	})
	assert.NoError(t, err)
	assert.Equal(t, ur.Token, "access")
	assert.Equal(t, ur.RefreshToken, "refresh")
	err = ur.FromProtoSignUp(&intrvproto.SignUpUserResponse{
		Id:                1,
		Username:          "username",
		VerificationToken: "abcd",
		CreatedAt:         nil,
	})
	assert.NoError(t, err)
	assert.Equal(t, ur.Username, "username")
	assert.Equal(t, ur.VerificationCode, "abcd")

	intrvU := &intrvproto.User{
		Id:        1,
		Username:  "username",
		Email:     "email",
		CreatedAt: nil,
	}
	err = ur.FromProtoUserResponse(&intrvproto.UserResponse{User: intrvU})
	assert.NoError(t, err)
	assert.Equal(t, ur.Username, "username")
}
