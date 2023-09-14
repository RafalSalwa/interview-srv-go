package router

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"os"

	gatewayConfig "github.com/RafalSalwa/interview-app-srv/cmd/gateway/config"
	"github.com/RafalSalwa/interview-app-srv/cmd/gateway/internal/middlewares"
	"github.com/RafalSalwa/interview-app-srv/docs"
	_ "github.com/RafalSalwa/interview-app-srv/docs"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NewApiRouter(cfg *gatewayConfig.Config, l *logger.Logger) *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.ContentTypeJson())
	router.Use(middlewares.CorrelationIDMiddleware())
	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.RequestLogMiddleware(l))

	setupHealthCheck(router)
	setupSwagger(router)

	if cfg.App.Debug {
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

func setupDebug(router *mux.Router) {
	router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
}

func setupSwagger(r *mux.Router) {
	docs.SwaggerInfo.Title = "Interview API for Gateway Service"
	docs.SwaggerInfo.Description = "API Gateway that works like a backends for frontends pattern and passes requests to specific services"

	jsonFile, err := os.Open("docs/swagger.json")
	if err != nil {
		fmt.Println(err)
	}
	bytesJSON, _ := io.ReadAll(jsonFile)
	docs.SwaggerInfo.SwaggerTemplate = string(bytesJSON)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods(http.MethodGet)
}
