package infra

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

type DBHandler struct {
	SqlConn   *sql.DB
	RedisConn *redis.Client
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

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	dbHander := new(DBHandler)
	dbHander.SqlConn = conn
	dbHander.RedisConn = redisClient
	return dbHander
}
