package jwt

import (
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func initConfig() config.ConfToken {
	c := config.New()
	return c.Token
}
func TestGenerateTokenPair(t *testing.T) {
	c := initConfig()

	issuedAt := time.Now()

	tp, err := GenerateTokenPair(c, 1, "testname")
	assert.NoError(t, err)

	at := tp.AccessToken
	token, err := ValidateToken(at, c.AccessTokenPublicKey)
	assert.NoError(t, err)
	expectedMap := map[string]interface{}{
		"user_id":  1.0,
		"username": "testname",
	}
	assert.NoError(t, err)
	assert.Equal(t, expectedMap, token)
	expiredAt := time.Now()
	require.WithinDuration(t, issuedAt, expiredAt, time.Second)

}

func TestCreateToken(t *testing.T) {
	c := initConfig()
	at, err := CreateToken(c.AccessTokenExpiresIn, "1", c.AccessTokenPrivateKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, at)
	if len(at) < 180 {
		t.Error("token not valid")
	}

}

func TestDecodeToken(t *testing.T) {
	c := initConfig()
	badToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	_, err := DecodeToken(badToken, c.AccessTokenPublicKey)
	assert.Error(t, err)
}
