package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	HTTPInfo  *HTTPInfo
	MySQLInfo *MySqlInfo
}

type HTTPInfo struct {
	Addr string
}

type MySqlInfo struct {
	MySQLUser     string
	MySQLPassword string
	MySQLAddr     string
	MySQLDBName   string
}

func LoadConfig() *appConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	addr := ":" + os.Getenv("PORT")

	httpInfo := &HTTPInfo{
		Addr: addr,
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlAddr := os.Getenv("MYSQL_ADDR")
	mysqlDBName := os.Getenv("MYSQL_DATABASE")

	dbInfo := &MySqlInfo{
		MySQLUser:     mysqlUser,
		MySQLPassword: mysqlPassword,
		MySQLAddr:     mysqlAddr,
		MySQLDBName:   mysqlDBName,
	}

	conf := appConfig{
		MySQLInfo: dbInfo,
		HTTPInfo:  httpInfo,
	}

	return &conf
}
