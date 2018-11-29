package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/go-xorm/xorm"
	"time"
)

var (
	SqlDB *sql.DB
	Dtsc *xorm.Engine
	Auth *xorm.Engine
)
func init() {
	init_dtsc()
	init_auth()
}

func init_dtsc() {
	connectionstring := "zh:******@(127.0.0.1:3306)/zh_dtsc?charset=utf8"
	maxidleconns := 5
	maxopenconns :=100
	db_conn, err := xorm.NewEngine("mysql", connectionstring)
	if err != nil {
		log.Fatalf("【dtsc.NewEngine】ex:%s\n", err.Error())
		return
	}
	err = db_conn.Ping()
	if err != nil {
		log.Fatalf("【dtsc.Ping】ex:%s\n", err.Error())
		return
	}
	db_conn.TZLocation = time.Local
	db_conn.SetMaxIdleConns(maxidleconns)
	db_conn.SetMaxOpenConns(maxopenconns)
	Dtsc = db_conn
}

func init_auth() {
	connectionstring := "zh:******@(127.0.0.1:3306)/zh_auth?charset=utf8"
	maxidleconns := 5
	maxopenconns :=100
	db_conn, err := xorm.NewEngine("mysql", connectionstring)
	if err != nil {
		log.Fatalf("【auth.NewEngine】ex:%s\n", err.Error())
		return
	}
	err = db_conn.Ping()
	if err != nil {
		log.Fatalf("【auth.Ping】ex:%s\n", err.Error())
		return
	}
	db_conn.TZLocation = time.Local
	db_conn.SetMaxIdleConns(maxidleconns)
	db_conn.SetMaxOpenConns(maxopenconns)
	Auth = db_conn
}
