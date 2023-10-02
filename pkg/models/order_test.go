package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderDBModel_TableName(t *testing.T) {
	o := OrderDBModel{}
	assert.Equal(t, "orders", o.TableName())
}
