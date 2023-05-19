package auth

import (
	apiHandler "github.com/RafalSalwa/interview-app-srv/api/handler"
	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
	"github.com/RafalSalwa/interview-app-srv/config"
	"net/http"
)

func (a *apiKeyAuth) middleware(h apiHandler.HandlerFunc) http.HandlerFunc {

	apiKeyHeader := a.cfg.APIKey

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//apiKey, err := bearerToken(r, apiKeyHeader)
			key := r.Header.Get("x-api-key")
			if key == "" {
				responses.NewUnauthorizedErrorResponse("API key missing")
				return
			}
			if apiKeyHeader != key {
				responses.NewBadRequestErrorResponse("wrong API key")
			}

			h(w, r)
		})
	}
}

type apiKeyAuth struct {
	h   apiHandler.HandlerFunc
	cfg config.ConfServer
}

func newApiKeyMiddleware(h apiHandler.HandlerFunc, cfg config.ConfServer) func(handler http.Handler) *apiKeyAuth {
	return &apiKeyAuth{
		h:   h,
		cfg: cfg,
	}

}
