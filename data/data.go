package data

import (
	"database/sql"
	"log"
	"time"
	"github.com/nu7hatch/gouuid"
	"fmt"
	"crypto/sha1"
	_ "github.com/lib/pq"
)

// should this be a global variable?
var Db * sql.DB

const(
	timeFormat = "Jan 2, 2006 at 3:04pm"
)

func init(){
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	return
}

func formatTimeToString(t time.Time) string{
	return t.Format(timeFormat)
}

func createUUID() (string,error){
	u4, err := uuid.NewV4()
	if err != nil{
		return "", fmt.Errorf("error while creating UUID:%s",err.Error())
	}
	return u4.String(), nil
}

func Encrypt(str string) string{
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}