package service

import (
	"fmt"

	"github.com/go-redis/redis"
)

const (
	EventTopic = "random_topic"
)

func ConnectToRedis() (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use the default DB
	})

	// Ping the Redis server to check if it's reachable.
	pong, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return nil, err
	}
	fmt.Println("Connected to Redis:", pong)

	return rdb, nil

}
