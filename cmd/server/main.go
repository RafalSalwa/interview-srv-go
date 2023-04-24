package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"

	unoHttp "github.com/RafalSalwa/interview/http"
	"github.com/RafalSalwa/interview/service"
	"github.com/RafalSalwa/interview/sql"
	"github.com/RafalSalwa/interview/util/logger"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func main() {
	log := logrus.New()
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})

	log.Info("just info message")
	// Output: Jan _2 15:04:05.000 [INFO] just info message

	log.WithField("component", "rest").Warn("warn message")
	// Output: Jan _2 15:04:05.000 [WARN] [rest] warn message
	env := os.Getenv("APP_ENV")

	usersDb := sql.NewUsersDB()

	us := service.NewMySqlService(usersDb)

	r := mux.NewRouter()

	handler := unoHttp.NewHandler(r, us)
	router := unoHttp.NewRouter(handler)
	server := unoHttp.NewServer(router)

	logger.Log(fmt.Sprintf("Server started - listen on address %s \n", server.Addr), logger.Info)
	unoHttp.SetupServer(server, env)
}
