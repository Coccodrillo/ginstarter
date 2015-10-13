package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

func GetBase() *Base {
	base := &Base{}
	db, err := sql.Open("mysql", os.Getenv("DB_STRING"))
	checkErr(err)
	db.SetMaxOpenConns(1400)
	base.Db = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	base.Db.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	return base
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
