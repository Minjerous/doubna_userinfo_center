package tool

import (
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

//redis 配置链接
//func init() {
//	redisCfg := GetCfg().Redis
//	RedisClient = redis.NewClient(&redis.Options{
//		Addr:     redisCfg.Addr + ":" + redisCfg.Port,
//		Password: redisCfg.Password,
//		DB:       redisCfg.Db,
//	})
//
//	ctx := context.Background()
//	ping, err := RedisClient.Ping(ctx).Result()
//	if err != nil {
//		fmt.Println("redis链接失败：", err.Error())
//		return
//	}
//	log.Println("redis链接成功", ping)
//}
