package database_mysql

import (
	"fmt"
	"net/url"
	"time"

	"github.com/jmoiron/sqlx"
)

func NewDB() *sqlx.DB {
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := "1234lupa"
	dbName := "test-db"

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
