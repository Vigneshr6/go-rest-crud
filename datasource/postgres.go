package datasource

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var conn *sqlx.DB

func init() {
	var err error
	conn, err = sqlx.Connect("postgres", "user=postgres password='postgres'")
	if err != nil {
		panic("db connection error")
	}
	conn.SetConnMaxIdleTime(2)
	conn.SetMaxOpenConns(5)
}

func GetConn() *sqlx.DB {
	return conn
}
