package router

import (
	"github.com/RafalSalwa/interview-app-srv/api/handler"
	"github.com/RafalSalwa/interview-app-srv/api/resource/health"
	"github.com/RafalSalwa/interview-app-srv/util/auth"
	_ "github.com/RafalSalwa/interview-app-srv/util/auth"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/api/resource/swagger"
	"github.com/gorilla/mux"
)

func NewRouter(handler handler.Handler, validator *validator.Validate) http.Handler {
	router := mux.NewRouter()

	health.SetupHealthCheck(router)
	swagger.SetupSwagger(router)

	var api = router.PathPrefix("/api").PathPrefix("/v1").Subrouter()

	setupUserRoutes(api, handler)
	return router
}

func setupUserRoutes(r *mux.Router, h handler.Handler) {
	r.Methods(http.MethodGet).Path("/users/{id}").HandlerFunc(auth.BasicAuth(h.GetUser()))
	r.Methods(http.MethodPost).Path("/users/{id}").HandlerFunc(auth.BasicAuth(h.PostUser()))
	r.Methods(http.MethodPost).Path("/users/change_password").HandlerFunc(auth.BasicAuth(h.PasswordChange()))
	r.Methods(http.MethodPost).Path("/users/auth").HandlerFunc(h.LogIn())
	r.Methods(http.MethodPut).Path("/users/registration").HandlerFunc(auth.BasicAuth(h.Create()))
	r.Methods(http.MethodPost).Path("/users/exist").HandlerFunc(auth.BasicAuth(h.UserExist()))
}
