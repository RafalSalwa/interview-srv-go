package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {
	os.Setenv("APP_ENV", "staging")
	path, err := getEnvPath()
	assert.NoError(t, err)
	assert.NotEmpty(t, path)
	assert.Contains(t, path, "gateway")

	c, err := InitConfig()
	assert.NoError(t, err)
	assert.NotEmpty(t, c.Http)
	assert.NotEmpty(t, c.Grpc)
	assert.NotEmpty(t, c.Auth)
	assert.NotEmpty(t, c.App)
	assert.NotEmpty(t, c.Logger)
	assert.NotEmpty(t, c.Jaeger)
	assert.NotEmpty(t, c.ServiceName)
	assert.Equal(t, c.ServiceName, "api_gateway_service")
}
