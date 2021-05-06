package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"kaishan/core/handlers/conf"
)

// Clients redis 客户端
var Clients map[string]*redis.Client

// InitRedis 初始化redis
func InitRedis(){
	configs := conf.Viper.GetStringMap("redis")
	Clients = make(map[string]*redis.Client, len(configs))
	for key, _ := range configs {
		client, err := newRedisClient(key)
		if err != nil {
			panic(fmt.Errorf("fatal err connect redis: %s", key))
		}

		Clients[key] = client
	}
}

func newRedisClient(key string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:       conf.Viper.GetString(fmt.Sprint("redis.", key, ".addr")),
		Password:   conf.Viper.GetString(fmt.Sprint("redis.", key, ".auth")),
		DB:         conf.Viper.GetInt(fmt.Sprint("redis.", key, ".db")),
		MaxRetries: 1,
	})

	_, err := client.Ping().Result()

	return  client, err
}

// GetClient 获取对应资源具柄的
func GetClient(name string) (client *redis.Client, err error) {
	client, ok := Clients[name]
	if !ok {
		err = errors.Errorf("fatal err get redis instance: %s", name)
	}
	return
}

// Close 关闭所有的具柄
func Close()  {
	for _, client := range Clients {
		client.Close()
	}
}