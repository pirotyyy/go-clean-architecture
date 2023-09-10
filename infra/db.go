package infra

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func SqlConnector() *SqlHandler {
	dbDriver := "mysql"
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbAddr := os.Getenv("MYSQL_ADDR")
	conn, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@(127.0.0.1:"+dbAddr+")/"+dbName)
	if err != nil {
		log.Println(err)
	}

	sqlHander := new(SqlHandler)
	sqlHander.Conn = conn
	return sqlHander
}
