package auth

import (
	"github.com/RafalSalwa/interview-app-srv/pkg/responses"
	"net/http"
)

type apiKeyAuth struct {
	apiKey string
}

func (a *apiKeyAuth) Middleware(h http.HandlerFunc) http.HandlerFunc {
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

func newApiKeyMiddleware(apiKey string) *apiKeyAuth {
	return &apiKeyAuth{
		apiKey: apiKey,
	}
}
