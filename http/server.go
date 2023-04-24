package http

import (
	"net/http"
	"os"
	"time"

	"github.com/RafalSalwa/interview/util/logger"
)

func NewServer(handler http.Handler) *http.Server {

	return &http.Server{
		Addr:         os.Getenv("LISTEN"),
		Handler:      handler,
		WriteTimeout: time.Duration(30) * time.Second,
		ReadTimeout:  time.Duration(30) * time.Second,
	}
}

func SetupServer(server *http.Server, env string) {
	var err error

	if env == "production" || env == "staging" {
		err = server.ListenAndServeTLS("/etc/ssl/private/unohouse.com.pl.crt", "/etc/ssl/private/unohouse.com.pl.key")
	} else {
		err = server.ListenAndServe()
	}

	if err != nil {
		logger.Log(err.Error(), logger.Error)
	}
}
