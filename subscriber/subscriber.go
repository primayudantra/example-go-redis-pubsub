package main

import (
	"fmt"

	"go-redis-pubsub-fx/service"
)

func main() {

	// Connect to Redis.
	rdb, err := service.ConnectToRedis()
	if err != nil {
		panic(err)
	}

	// Subscribe to a channel.
	pubsub := rdb.Subscribe(service.EventTopic)
	defer pubsub.Close()

	// Handle incoming messages from Redis Pub/Sub.
	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			fmt.Println("Error receiving message:", err)
			return
		}
		fmt.Println("Received message from Redis Pub/Sub:", msg.Payload)
	}

}
