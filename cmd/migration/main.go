package main

import (
	"embed"
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"github.com/RafalSalwa/interview-app-srv/sql"
	"github.com/pressly/goose/v3"
	"log"
)

var embedMigrations embed.FS

func main() {
	c := config.New()
	l := logger.NewConsole(c.App.Debug)
	db := sql.NewUsersDB(c.DB, l)
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf(err.Error())
		}
	}()
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	if err := goose.Up(db.DB, "migrations"); err != nil {
		panic(err)
	}

}
