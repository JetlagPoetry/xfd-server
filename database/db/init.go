package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
	"xfd-backend/config"
)

var (
	mySQL *gorm.DB
)

func NewMySQL() {
	mysqlDB, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.GetString("mysql.user"),
		config.Cfg.GetString("mysql.password"),
		config.Cfg.GetString("mysql.host"),
		config.Cfg.GetString("mysql.name"))), &gorm.Config{})
	if err != nil {
		panic("Init MySQL failed")
	}

	sqlDB, _ := mysqlDB.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	mySQL = mysqlDB
	log.Println("MySQL init success")
}
