package redis

import "time"

func Lock(key string, expiration time.Duration) bool {
	ok, err := RedisClient.SetNX(key, 1, expiration).Result()
	if err != nil {
		return false
	}
	return ok
}

func Unlock(key string) error {
	_, err := RedisClient.Del(key).Result()
	return err
}
