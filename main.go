package main

import (
	"fmt"
	"log"
	"net/http"
	"xfd-backend/database/db"
	"xfd-backend/router"
)

func main() {
	Init()

	//cron.StartCron()

	r := router.NewRouter()
	log.Println("==================Server Start===================")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func Init() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("Mysql init failed with %+v", err))
	}
	//config.InitConfig()
}
