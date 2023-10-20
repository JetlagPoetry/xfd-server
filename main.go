package main

import (
	"fmt"
	"log"
	"net/http"
	"xfd-backend/database/db"
	"xfd-backend/service"
)

func main() {
	Init()

	//r := router.NewRouter()
	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func Init() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("Mysql init failed with %+v", err))
	}
	//config.InitConfig()
}
