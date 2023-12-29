package cache

import (
	"fmt"
	"recruit/settings"

	"github.com/go-redis/redis"
)

var RedisCache = &redis.Client{}

func Init(rds *settings.RedisConfig) error {
	address := fmt.Sprintf("%s:%d", rds.Host, rds.Port)
	RedisCache = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: rds.Password,
		DB:       rds.DB,
	})
	//ping
	pong, err := RedisCache.Ping().Result()
	if err != nil {
		fmt.Println("ping error", err.Error())
		return err
	}
	fmt.Println("ping result:", pong)
	return nil
}

func Close() {
	_ = RedisCache.Close()
}
