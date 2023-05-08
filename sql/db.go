package sql

import (
	"database/sql"
	"fmt"
	"github.com/RafalSalwa/interview-app-srv/config"
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

const (
	driver   = "mysql"
	dbString = "%s:%s@tcp(%s:%d)/%s?%s"
	dbParams = "parseTime=true&loc=Europe%2FWarsaw&charset=utf8&collation=utf8_polish_ci"
)

func NewUsersDB(c config.ConfDB, l *logger.Logger) DB {
	con := fmt.Sprintf(dbString, c.Username, c.Password, c.Host, c.Port, c.DBName, dbParams)
	db, err := sql.Open(driver, con)
	if err != nil {
		l.Error().Err(err)

	}
	err = db.Ping()
	if err != nil {
		l.Error().Err(err)

	}
	return DB{db}
}
