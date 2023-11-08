package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

var dbInstance *gorm.DB

// Init 初始化数据库
func Init() error {

	source := "%s:%s@tcp(%s)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true"
	user := os.Getenv("MYSQL_USERNAME")
	pwd := os.Getenv("MYSQL_PASSWORD")
	addr := os.Getenv("MYSQL_ADDRESS")
	dataBase := os.Getenv("MYSQL_DATABASE")
	if dataBase == "" {
		dataBase = "golang_demo"
	}
	source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	fmt.Println("start init mysql with ", source)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		}})
	if err != nil {
		fmt.Println("DB Open error,err=", err.Error())
		return err
	}
	db = db.Debug()
	//_ = db.AutoMigrate(&model.Category{})
	//_ = db.AutoMigrate(&model.AreaCode{})
	//_ = db.AutoMigrate(&model.Goods{})
	//_ = db.AutoMigrate(&model.Specification{})
	//_ = db.AutoMigrate(&model.SpecificationValue{})
	//_ = db.AutoMigrate(&model.ProductVariant{})

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("DB Init error,err=", err.Error())
		return err
	}

	// 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(15)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(250)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbInstance = db

	fmt.Println("finish init mysql with ", source)
	return nil
}

// Get ...
func Get() *gorm.DB {
	return dbInstance
}
