package auth

import (
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	apiHandler "github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/handler"
)

type apiKeyAuth struct {
	h      apiHandler.AuthHandler
	apiKey string
}

func (a *apiKeyAuth) Middleware(h apiHandler.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("x-api-key")
		if key == "" {
			responses.NewUnauthorizedErrorResponse("API key missing")
			return
		}
		if a.apiKey != key {
			responses.NewBadRequestErrorResponse("wrong API key")
		}

		h(w, r)
	}
}

func newApiKeyMiddleware(h apiHandler.AuthHandler, apiKey string) *apiKeyAuth {
	return &apiKeyAuth{
		h:      h,
		apiKey: apiKey,
	}
}
