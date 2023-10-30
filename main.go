package main

import (
	"fmt"
	"log"
	"net/http"
	"xfd-backend/database/db"
	"xfd-backend/pkg/cron"
	"xfd-backend/router"
)

func main() {
	Init()

	cron.StartCron()

	r := router.NewRouter()
	log.Println("==================Server Start===================")
	log.Fatal(http.ListenAndServe(":80", r))

}

func Init() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("Mysql init failed with %+v", err))
	}
	//config.InitConfig()
}
