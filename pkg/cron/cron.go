package cron

import (
	"context"
	"github.com/robfig/cron/v3"
	"log"
	"xfd-backend/service"
)

func StartCron() {
	c := cron.New()
	c.AddFunc("* 0 * * * ?", func() {
		log.Println("[Cron] ProcessPointVerify start")
		// todo 分布式锁
		err := service.NewOrgService().ProcessPointVerify(context.Background())
		if err != nil {
			log.Println("[Cron] ProcessPointVerify failed, err=", err)
			return
		}
		log.Println("[Cron] ProcessPointVerify success")
	})

	c.AddFunc("* 5 * * * ?", func() {
		log.Println("[Cron] ProcessPointDistribute start")
		// todo 分布式锁
		err := service.NewOrgService().ProcessPointDistribute(context.Background())
		if err != nil {
			log.Println("[Cron] ProcessPointDistribute failed, err=", err)
			return
		}
		log.Println("[Cron] ProcessPointDistribute success")
	})

	c.AddFunc("* * 0 * * ?", func() {
		log.Println("[Cron] ProcessPointDistribute start")
		// todo 分布式锁
		err := service.NewOrgService().ProcessPointDistribute(context.Background())
		if err != nil {
			log.Println("[Cron] ProcessPointDistribute failed, err=", err)
			return
		}
		log.Println("[Cron] ProcessPointDistribute success")
	})

	c.AddFunc("*/10 * * * * ?", func() {
		log.Println("[Cron] SetCategoryCache start")
		err := service.NewMallService().SetCategoryCache(context.Background())
		if err != nil {
			log.Println("[Cron] SetCategoryCache failed, err=", err)
			return
		}
		log.Println("[Cron] SetCategoryCache success")
	})

	c.Start()
}
