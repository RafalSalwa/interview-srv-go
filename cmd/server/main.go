package main

import (
	"fmt"
	"os"

	unoHttp "github.com/RafalSalwa/interview/http"
	"github.com/RafalSalwa/interview/service"
	"github.com/RafalSalwa/interview/sql"
	"github.com/RafalSalwa/interview/tools/logger"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"field1": "foo",
			"field2": "bar",
		},
	).Info("Log message here!!!")

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
