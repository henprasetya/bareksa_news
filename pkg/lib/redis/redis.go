package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Redis struct {
	Redis *redis.Client
}

func NewRedis() *Redis {
	s := connectToRedis()

	return s
}

func connectToRedis() *Redis {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // default DB
	})
	result, err := redisPing(redisClient)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	return &Redis{
		Redis: redisClient,
	}
}

func redisPing(client *redis.Client) (string, error) {
	result, err := client.Ping().Result()
	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}
