package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		log.Fatal("DB close error: ", err)
	}
}

func init() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", user, password, database)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB open error: ", err.Error())
	}

	checkConnection(10)
	fmt.Println("DB connected")
}

func checkConnection(count int) {
	err := db.Ping()
	if err == nil {
		return
	}

	if count > 0 {
		time.Sleep(time.Second)
		count--
		fmt.Printf("retry... %v\n", count)
		checkConnection(count)
	} else {
		log.Fatal("DB connection error: ", err.Error())
	}
}
