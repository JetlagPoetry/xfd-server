package main

import (
	"fmt"
	"log"
	"net/http"
	"xfd-backend/config"
	"xfd-backend/database/db"
	"xfd-backend/router"
)

func main() {
	Init()

	r := router.NewRouter()
	//http.HandleFunc("/", service.IndexHandler)
	//http.HandleFunc("/api/count", service.CounterHandler)
	log.Fatal(http.ListenAndServe(":80", r))
}

func Init() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("Mysql init failed with %+v", err))
	}
	config.InitConfig()
}
