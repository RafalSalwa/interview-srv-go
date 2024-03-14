package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubscriptionDBModel_TableName(t *testing.T) {
	p := SubscriptionDBModel{}
	assert.Equal(t, "subscription", p.TableName())
}
