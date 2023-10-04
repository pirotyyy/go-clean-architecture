package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBHandler struct {
	SqlConn *sql.DB
}

func DBConnector() *DBHandler {
	dbDriver := "mysql"
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbAddr := os.Getenv("MYSQL_ADDR")
	conn, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@(127.0.0.1:"+dbAddr+")/"+dbName)
	if err != nil {
		log.Println(err)
	}

	dbHander := new(DBHandler)
	dbHander.SqlConn = conn
	return dbHander
}
