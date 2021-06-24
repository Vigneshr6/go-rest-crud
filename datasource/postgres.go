package datasource

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var conn *sqlx.DB

func init() {
	log.Info("Stage : " + os.Getenv("stage"))
	var err error
	conn, err = sqlx.Connect("postgres", fmt.Sprintf("user=%s password='%s'", os.Getenv("user"), os.Getenv("password")))
	if err != nil {
		panic("db connection error")
	}
	conn.SetConnMaxIdleTime(2)
	conn.SetMaxOpenConns(5)
}

func GetConn() *sqlx.DB {
	return conn
}
