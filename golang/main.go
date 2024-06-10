package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	var conn *amqp.Connection
	var err error

	for conn == nil {
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err != nil {
			log.Printf("Failed to connect to RabbitMQ. Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
		}
	}

	defer conn.Close()

	log.Println("Successfully connected to RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"golang", // Queue
		false,    // Durable
		false,    // Delete when unused
		false,    // Exclusive
		false,    // No-wait
		nil,      // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer tag
		true,   // Auto-acknowledgement
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received: %s", d.Body)
		}
	}()

	log.Println("Waiting for messages.")
	<-forever
}
