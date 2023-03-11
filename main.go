package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/ymho/qnin-sp/api"
)

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	r := api.NewRouter(db)

	log.Println("server start at port 80")
	log.Fatal(http.ListenAndServe(":80", r))
}
