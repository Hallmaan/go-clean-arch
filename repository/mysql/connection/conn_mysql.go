package connection

import (
	"fmt"
	"net/url"
	"time"

	"github.com/jmoiron/sqlx"
)

func NewDB(host string, port string, username string, pass string, dbname string) *sqlx.DB {
	dbHost := host
	dbPort := port
	dbUser := username
	dbPass := pass
	dbName := dbname

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, err := sqlx.Open(`mysql`, dsn)

	if err != nil {
		panic(err)
	}

	err = dbConn.Ping()

	dbConn.SetMaxIdleConns(5)
	dbConn.SetMaxOpenConns(20)
	dbConn.SetConnMaxLifetime(60 * time.Minute)
	dbConn.SetConnMaxIdleTime(10 * time.Minute)

	return dbConn
}
