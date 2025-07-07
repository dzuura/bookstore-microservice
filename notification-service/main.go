package main

import (
	"log"
	"github.com/dzuura/bookstore-microservice/notification-service/subscriber"
)

func main() {
	log.Println("Notification service is starting...")
	err := subscriber.SubscribeToOrderPlaced()
	if err != nil {
		log.Fatal("Failed to subscribe to NATS:", err)
	}
}
