package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"xfd-backend/database/db"
	"xfd-backend/pkg/cron"
	"xfd-backend/pkg/jwt"
	"xfd-backend/pkg/utils"
	"xfd-backend/router"
	"xfd-backend/service"
)

func main() {
	Init()

	go cron.StartCron()

	r := router.NewRouter()
	log.Println("==================Server Start===================")
	log.Fatal(http.ListenAndServe(":80", r))

}

func Init() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("Mysql init failed with %+v", err))
	}
	utils.InitSms()
	jwt.Init()
	//config.InitConfig()
	if err := initCache(); err != nil {
		panic(fmt.Sprintf("Local cache init failed with %+v", err))
	}
}

func initCache() error {
	err := service.NewMallService().SetCategoryCache(context.Background())
	if err != nil {
		return err
	}
	return nil
}
