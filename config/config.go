package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	Cfg *viper.Viper
)

func InitConfig() {
	vip := viper.New()
	vip.AddConfigPath("./config/")
	vip.SetConfigName("config")
	vip.SetConfigType("toml")
	err := vip.ReadInConfig()
	if err != nil {
		panic("Init config failed")
	}
	Cfg = vip
	log.Println("Config init success")
}
