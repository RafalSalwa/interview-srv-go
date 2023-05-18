package router

import (
	"encoding/json"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/api/resource/middlewares"
	"github.com/RafalSalwa/interview-app-srv/config"
	"net/http"

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

func NewApiRouter(l *logger.Logger, c config.ConfToken) *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.ContentTypeJson())
	router.Use(middlewares.CorrelationIDMiddleware())
	//router.Use(middlewares.RequestLogMiddleware(l))
	router.Use(middlewares.CorsMiddleware())
	//router.Use(middlewares.DeserializeUser(c))
	setupIndexPageRoutesInfo(router)
	setupHealthCheck(router)

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

func GetRoutesList(r *mux.Router) []AppRoute {
	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		routePath, err := route.GetPathTemplate()
		if err != nil {

		}
		routeMethods, err := route.GetMethods()
		if err != nil {

		}
		appRoutes = append(appRoutes, AppRoute{Path: routePath, Methods: routeMethods})
		return nil
	})
	if err != nil {
		return nil
	}
	for _, r := range appRoutes {
		fmt.Printf("'%-8s'", r.Methods)
		fmt.Println(r.Path)
	}

	return appRoutes
}