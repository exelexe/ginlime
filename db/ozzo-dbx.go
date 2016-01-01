package db

import (
	"fmt"
	"github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitDb() *dbx.DB {
	log.Println("InitDb called.")
	var datasource string
	params := "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	datasource = fmt.Sprintf("root:root@tcp(localhost:3306)/demo%s", params)

	db, _ := dbx.Open("mysql", datasource)
	db.LogFunc = log.Printf
	return db
}
