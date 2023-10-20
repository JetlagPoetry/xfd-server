package main

import (
	"log"
	"net/http"
	"xfd-backend/config"
	"xfd-backend/router"
)

func main() {
	Init()
	r := router.NewRouter()
	server := &http.Server{
		Addr:    "127.0.0.1:60010",
		Handler: r,
	}
	log.Println("============== XFD-Backend Server Start ==============")
	_ = server.ListenAndServe()

}

func Init() {
	config.InitConfig()
	//db.NewMySQL()
}
