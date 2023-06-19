package health

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupHealthCheck(r *mux.Router) {
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ok"))
	}).Methods(http.MethodGet)
}
