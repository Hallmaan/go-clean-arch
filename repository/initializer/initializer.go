package mysql_initializer

import (
	database_mysql "clean_arch_ws/repository/mysql"
	"github.com/jmoiron/sqlx"
)

type Replication struct {
	Primary *sqlx.DB
	Standby *sqlx.DB
}

type MysqlInitializer struct {
	DbHostPrimary string
	DbHostStandby string
	DbPort        string
	DbUser        string
	DbPass        string
	DbName        string
}

func (mi MysqlInitializer) MySQLInit() *Replication {
	return &Replication{
		Primary: database_mysql.NewDB(mi.DbHostPrimary, mi.DbPort, mi.DbUser, mi.DbPass, mi.DbName),
		Standby: database_mysql.NewDB(mi.DbHostPrimary, mi.DbPort, mi.DbUser, mi.DbPass, mi.DbName),
	}
}
