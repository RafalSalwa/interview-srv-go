package http

import (
	"encoding/json"
	"net/http"

	"github.com/RafalSalwa/interview/swagger"
	"github.com/RafalSalwa/interview/util/logger"
	"github.com/gorilla/mux"
)

type AppRoute struct {
	Path    string   `json:"path"`
	Methods []string `json:"methods"`
}

var appRoutes []AppRoute

func NewRouter(handler Handler) http.Handler {
	router := mux.NewRouter()

	setupHealthCheck(router)
	setupSwagger(router)
	setupIndexPageRoutesInfo(router)
	setupUserRoutes(router, handler)
	getRoutesList(router)
	return router
}

func setupHealthCheck(r *mux.Router) {
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Up"))
	}).Methods(http.MethodGet)
}

func setupIndexPageRoutesInfo(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		js, err := json.MarshalIndent(appRoutes, "", "    ")
		if err != nil {
			logger.Log(err.Error(), logger.Error)
			respondInternalServerError(w)
		}

		respond(w, http.StatusOK, js)
	}).Methods(http.MethodGet)

}

func setupSwagger(r *mux.Router) {
	h := http.FileServer(http.FS(swagger.GetStaticFiles()))
	r.PathPrefix("/swagger").Handler(h).Methods(http.MethodGet)
}

func setupUserRoutes(r *mux.Router, h Handler) {
	r.Methods(http.MethodGet).Path("/users/{id}").HandlerFunc(BasicAuth(h.getUserById()))
	r.Methods(http.MethodPost).Path("/users/{id}").HandlerFunc(BasicAuth(h.postUsersIdEdit()))
	r.Methods(http.MethodPost).Path("/users/change_password").HandlerFunc(BasicAuth(h.postUsersIdChangePassword()))
	r.Methods(http.MethodPost).Path("/users/auth").HandlerFunc(h.postUsersLogin())
	r.Methods(http.MethodPost).Path("/users/registration").HandlerFunc(BasicAuth(h.postUsersRegistration()))
	r.Methods(http.MethodPost).Path("/users/exist").HandlerFunc(BasicAuth(h.postUsersExist()))
}

func getRoutesList(router *mux.Router) []AppRoute {
	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		routePath, err := route.GetPathTemplate()
		if err != nil {
			logger.Log(err.Error(), logger.Error)
		}
		routeMethods, err := route.GetMethods()
		if err != nil {
			logger.Log(err.Error(), logger.Error)
		}
		appRoutes = append(appRoutes, AppRoute{Path: routePath, Methods: routeMethods})
		return nil
	})
	if err != nil {
		return nil
	}
	return appRoutes
}
