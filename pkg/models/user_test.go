package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModels(t *testing.T) {
	dbu := UserDBModel{
		Id: 1,
	}
	assert.NotEmpty(t, dbu)
	err := dbu.BeforeCreate(nil)
	assert.NoError(t, err)
	assert.Equal(t, false, dbu.Active)
	assert.Equal(t, "user", dbu.TableName())
}
