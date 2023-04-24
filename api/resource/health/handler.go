package health

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupHealthCheck(r *mux.Router) {
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Ok"))
	}).Methods(http.MethodGet)
}
