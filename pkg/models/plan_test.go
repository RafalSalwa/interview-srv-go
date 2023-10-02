package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlan_TableName(t *testing.T) {
	p := Plan{}
	assert.Equal(t, "plan", p.TableName())
}
