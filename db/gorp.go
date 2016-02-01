package db
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"gopkg.in/gorp.v1"
//	"log"
//	"os"
//)
//
//func InitDbMap() *gorp.DbMap {
//	var datasource string
//	params := "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
//
//	// TODO config
//	datasource = fmt.Sprintf("root:root@tcp(localhost:3306)/demo%s", params)
//
//	// TODO error handling
//	db, err := sql.Open("mysql", datasource)
//	if err != nil {
//		log.Println(err)
//	}
//
//	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
//	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "[SQL] ", log.Lmicroseconds))
//	return dbmap
//}
