package data

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:147369@tcp(140.143.239.161:3306)/goapi?parseTime=true")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	return
}