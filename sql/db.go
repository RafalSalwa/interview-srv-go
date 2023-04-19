package sql

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func NewDB() DB {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@"+"tcp("+os.Getenv("MYSQL_HOSTS")+":"+os.Getenv("MYSQL_PORT")+")/"+os.Getenv("MYSQL_MAIN_DATABASE")+"?parseTime=true&loc=Europe%2FWarsaw&charset=utf8&collation=utf8_polish_ci")
	if err != nil {
		log.Fatalln("unable to connect to mySQL", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln("unable to connect to mySQL", err)
	}

	return DB{db}
}

func NewUsersDB() DB {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@"+"tcp("+os.Getenv("MYSQL_HOSTS")+":"+os.Getenv("MYSQL_PORT")+")/"+os.Getenv("MYSQL_USERS_DATABASE")+"?parseTime=true&loc=Europe%2FWarsaw&charset=utf8&collation=utf8_polish_ci")
	if err != nil {
		log.Fatalln("unable to connect to mySQL", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln("unable to connect to mySQL", err)
	}

	return DB{db}
}
