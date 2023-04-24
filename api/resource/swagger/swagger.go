package swagger

import (
	"embed"
	"github.com/gorilla/mux"
	"net/http"
)

var static embed.FS

func GetStaticFiles() embed.FS {
	return static
}

func SetupSwagger(r *mux.Router) {
	h := http.FileServer(http.FS(GetStaticFiles()))
	r.PathPrefix("/swagger").Handler(h).Methods(http.MethodGet)
}
