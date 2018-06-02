package data

import (
	"database/sql"
	"log"
)

// should this be a global variable?
var Db * sql.DB

func init(){
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	return
}
