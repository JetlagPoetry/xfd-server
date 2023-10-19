package main

import (
	"log"
	"xfd-backend/config"
	"xfd-backend/database/db"
	"xfd-backend/router"
)

func main() {
	Init()
	r := router.NewRouter()
	r.Run("127.0.0.1:60010")
	log.Println("============== XFD-Backend Server Start ==============")
}

func Init() {
	config.InitConfig()
	db.NewMySQL()
}
