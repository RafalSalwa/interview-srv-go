package swagger

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupSwagger(r *mux.Router) {
	sh := http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger/")))
	r.PathPrefix("/swagger/").Handler(sh)
}
