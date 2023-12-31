package cron

import (
	"context"
	"github.com/robfig/cron/v3"
	"log"
	"time"
	"xfd-backend/database/redis"
	"xfd-backend/service"
)

func StartCron() {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc("0 */5 * * * ?", func() {
		//log.Println("[Cron] ProcessPointVerify start")

		ok := redis.Lock("cron-process-point-verify", time.Minute*10)
		if !ok {
			return
		}
		defer redis.Unlock("cron-process-point-verify")

		err := service.NewOrgService().ProcessPointVerify(context.Background())
		if err != nil {
			log.Println("[Cron] ProcessPointVerify failed, err=", err)
			return
		}
		log.Println("[Cron] ProcessPointVerify success")
	})
	if err != nil {
		log.Println("[Cron] ProcessPointVerify failed, err=", err)
	}

	_, err = c.AddFunc("30 */3 * * * ?", func() {
		//log.Println("[Cron] ProcessPointDistribute start")

		ok := redis.Lock("cron-process-point-distribute", time.Minute*10)
		if !ok {
			return
		}
		defer redis.Unlock("cron-process-point-distribute")

		err := service.NewOrgService().ProcessPointDistribute(context.Background())
		if err != nil {
			log.Println("[Cron] ProcessPointDistribute failed, err=", err)
			return
		}
		log.Println("[Cron] ProcessPointDistribute success")
	})
	if err != nil {
		log.Println("[Cron] ProcessPointDistribute failed, err=", err)
	}

	_, err = c.AddFunc("0 0 * * * ?", func() {
		//log.Println("[Cron] ProcessPointExpired start")

		ok := redis.Lock("cron-process-point-expire", time.Minute*10)
		if !ok {
			return
		}
		defer redis.Unlock("cron-process-point-expire")

		err := service.NewOrgService().ProcessPointExpired(context.Background())
		if err != nil {
			log.Println("[Cron] ProcessPointExpired failed, err=", err)
			return
		}
		log.Println("[Cron] ProcessPointExpired success")
	})
	if err != nil {
		log.Println("[Cron] ProcessPointExpired failed, err=", err)
	}

	_, err = c.AddFunc("0 */6 * * * ?", func() {
		//log.Println("[Cron] SetCategoryCache start")

		ok := redis.Lock("cron-set-category-cache", time.Minute*10)
		if !ok {
			return
		}
		defer redis.Unlock("cron-set-category-cache")

		err := service.NewMallService().SetCategoryCache(context.Background())
		if err != nil {
			log.Println("[Cron] SetCategoryCache failed, err=", err)
			return
		}
		log.Println("[Cron] SetCategoryCache success")
	})
	if err != nil {
		log.Println("[Cron] SetCategoryCache failed, err=", err)
	}

	_, err = c.AddFunc("* * * * * ?", func() {
		//log.Println("[Cron] BatchPaymentLookup start")

		ok := redis.Lock("cron-set-payment-lookup", time.Minute*10)
		if !ok {
			return
		}
		defer redis.Unlock("cron-set-payment-lookup")

		err := service.NewOrderService().BatchPaymentLookup(context.Background())
		if err != nil {
			log.Println("[Cron] BatchPaymentLookup failed, err=", err)
			return
		}
		//log.Println("[Cron] BatchPaymentLookup success")
	})
	if err != nil {
		log.Println("[Cron] BatchPaymentLookup failed, err=", err)
	}

	_, err = c.AddFunc("*/30 * * * * ?", func() {
		//log.Println("[Cron] BatchPointConfirm start")

		ok := redis.Lock("cron-batch-point-confirm", time.Minute*10)
		if !ok {
			return
		}
		defer redis.Unlock("cron-batch-point-confirm")

		err := service.NewOrderService().BatchPointConfirm(context.Background())
		if err != nil {
			log.Println("[Cron] BatchPointConfirm failed, err=", err)
			return
		}
		log.Println("[Cron] BatchPointConfirm success")
	})
	if err != nil {
		log.Println("[Cron] BatchPointConfirm failed, err=", err)
	}

	_, err = c.AddFunc("0 */1 * * * ?", func() {
		ok := redis.Lock("cron--auto-order-confirm-receipt", time.Minute*10)
		if !ok {
			return
		}
		defer redis.Unlock("cron--auto-order-confirm-receipt")

		err := service.NewOrderService().AutoOrderConfirmReceipt(context.Background())
		if err != nil {
			log.Println("[Cron] AutoOrderConfirmReceipt failed, err=", err)
			return
		}
		log.Println("[Cron] AutoOrderConfirmReceipt success")
	})
	if err != nil {
		log.Println("[Cron] AutoOrderConfirmReceipt failed, err=", err)
	}
	c.Start()
}
