package initializer

import (
	conn "clean_arch_ws/repository/mysql/connection"
	"github.com/jmoiron/sqlx"
)

type Replication struct {
	Primary *sqlx.DB
	Standby *sqlx.DB
}

func NewMySQLInit() *Replication {
	return &Replication{
		Primary: conn.NewDB("localhost", "3306", "root", "1234lupa", "test-db"),
		Standby: conn.NewDB("localhost", "3306", "root", "1234lupa", "test-db"),
	}
}
