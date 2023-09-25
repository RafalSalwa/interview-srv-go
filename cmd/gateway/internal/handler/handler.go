package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type RouteRegisterer interface {
	RegisterRoutes(r *mux.Router, cfg interface{})
}
