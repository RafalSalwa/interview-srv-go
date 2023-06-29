package auth

import (
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	simpleHandler "github.com/RafalSalwa/interview-app-srv/pkg/simple_handler"
	"net/http"
)

type apiKeyAuth struct {
	h      simpleHandler.HandlerFunc
	apiKey string
}

func (a *apiKeyAuth) middleware(h simpleHandler.HandlerFunc) http.HandlerFunc {
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

func newApiKeyMiddleware(h simpleHandler.HandlerFunc, apiKey string) *apiKeyAuth {
	return &apiKeyAuth{
		h:      h,
		apiKey: apiKey,
	}
}
