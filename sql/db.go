package sql

import (
	"database/sql"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DB struct {
	*sql.DB
}

const (
	driver   = "mysql"
	dbString = "%s:%s@tcp(%s:%d)/%s?%s"
	dbParams = "parseTime=true&loc=Europe%2FWarsaw&charset=utf8&collation=utf8_polish_ci"
)

func NewUsersDB(c config.ConfDB) DB {
	con := fmt.Sprintf(dbString, c.Username, c.Password, c.Host, c.Port, c.DBName, dbParams)
	fmt.Println(con)
	db, err := sql.Open(driver, con)
	if err != nil {
		log.Fatalln("unable to connect to mySQL", err)
	}

	return DB{db}
}
