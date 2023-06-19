package router

import (
	"encoding/json"
	"net/http"

	"github.com/RafalSalwa/interview-app-srv/api/resource/middlewares"
	"github.com/RafalSalwa/interview-app-srv/config"

	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/gorilla/mux"

	"github.com/RafalSalwa/interview-app-srv/api/resource/responses"
)

type AppRouter struct {
	Router  *mux.Router
	Handler http.Handler
	Logger  *logger.Logger
}

type AppRoute struct {
	Path    string   `json:"path"`
	Methods []string `json:"methods"`
}

var appRoutes []AppRoute

func NewApiRouter(l *logger.Logger, c *config.Conf) *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.ContentTypeJson())
	router.Use(middlewares.CorrelationIDMiddleware())
	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.RequestLogMiddleware(l))

	setupIndexPageRoutesInfo(router)
	setupHealthCheck(router)

	if c.App.Debug {
		setupDebug(router)
	}

	return router
}

func setupHealthCheck(router *mux.Router) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Up"))
	}).Methods(http.MethodGet)
}

func setupIndexPageRoutesInfo(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		js, err := json.MarshalIndent(appRoutes, "", "    ")
		if err != nil {
			responses.RespondInternalServerError(w)
		}

		responses.Respond(w, http.StatusOK, js)
	}).Methods(http.MethodGet)
}

func setupDebug(router *mux.Router) {
	router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
}
