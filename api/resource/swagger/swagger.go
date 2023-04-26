package swagger

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupSwagger(r *mux.Router) {
	sh := http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger/")))
	r.PathPrefix("/swagger/").Handler(sh)
}
