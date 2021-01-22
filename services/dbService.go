package services

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type DbServiceType struct{}

var dbClient *sql.DB
var DbService DbServiceType

func init() {
	DbService = DbServiceType{}
	dbClient = DbService.InitClient()
}

func (*DbServiceType) InitClient() *sql.DB {
	host := os.Getenv("DB_HOST")
	if len(host) == 0 {
		log.Fatal("DB_HOST env variable not set")
	}
	user := os.Getenv("DB_USER")
	if len(user) == 0 {
		log.Fatal("DB_USER env variable not set")
	}
	pwd := os.Getenv("DB_PWD")
	if len(pwd) == 0 {
		log.Fatal("DB_PWD env variable not set")
	}
	port := os.Getenv("DB_PORT")
	if len(port) == 0 {
		log.Fatal("DB_PORT env variable not set")
	}
	database := os.Getenv("DB_DB")
	if len(database) == 0 {
		log.Fatal("DB_DB env variable not set")
	}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pwd, host, port, database))
	if err != nil {
		log.Fatal("connection to DB failed")
	}

	return db
}
func (*DbServiceType) Query(query string) (*sql.Rows, error) {
	res, err := dbClient.Query(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (*DbServiceType) Exec(query string) (sql.Result, error) {
	res, err := dbClient.Exec(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}
