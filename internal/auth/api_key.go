package auth

import (
	apiHandler "github.com/RafalSalwa/interview-app-srv/api/handler"
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"net/http"
)

type apiKeyAuth struct {
	h      apiHandler.HandlerFunc
	apiKey string
}

func (a *apiKeyAuth) middleware(h apiHandler.HandlerFunc) http.HandlerFunc {

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

func newApiKeyMiddleware(h apiHandler.HandlerFunc, apiKey string) *apiKeyAuth {
	return &apiKeyAuth{
		h:      h,
		apiKey: apiKey,
	}
}
