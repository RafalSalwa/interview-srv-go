package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaymentDBModel_TableName(t *testing.T) {
	p := PaymentDBModel{}
	assert.Equal(t, "payment", p.TableName())
}
