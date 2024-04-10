package main

import (
	"encoding/json"
	"fmt"
	"go-redis-pubsub-fx/service"
	"math/rand"
	"time"
)

var (
	exchangeRates = map[string]float64{
		"USD-SGD": 0.75, // Initial exchange rate
		"EUR-SGD": 0.65, // Initial exchange rate
		"GBP-SGD": 0.55, // Initial exchange rate
		"JPY-SGD": 82.0, // Initial exchange rate
	}
	currencies = []string{"USD-SGD", "EUR-SGD", "GBP-SGD", "JPY-SGD"}
)

func main() {
	// var ctx = context.Background()
	rdb, err := service.ConnectToRedis()
	if err != nil {
		panic(err)
	}
	for {
		// Generate random exchange rates for each currency.
		for _, currency := range currencies {
			exchangeRates[currency] = rand.Float64() * 10 // Generates a random rate between 0 and 10
		}
		fmt.Println("Updated exchange rates:", exchangeRates)

		s, _ := json.Marshal(exchangeRates)
		// Publish message to Redis Pub/Sub channel.
		err := rdb.Publish(service.EventTopic, string(s)).Err()
		if err != nil {
			fmt.Println("Error publishing message to Redis:", err)
			return
		}
		fmt.Println("Published message to Redis Pub/Sub:", string(s))

		time.Sleep(5 * time.Second) // Wait for 5 seconds before updating rates again
	}

}
