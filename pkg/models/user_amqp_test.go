package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserDBModel_AMQP(t *testing.T) {
	um := UserDBModel{
		Id:               1,
		Username:         "test",
		Email:            "test@test.com",
		VerificationCode: "abcdef",
	}
	m := um.AMQP()
	assert.NotEmpty(t, m)

}
