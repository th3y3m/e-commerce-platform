package RabbitMQ

import (
	"log"

	"github.com/streadway/amqp"
)

// Publish message to RabbitMQ
func PublishMessage(message string) error {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declare a queue (create it if it doesn't exist)
	queue, err := ch.QueueDeclare(
		"order_queue", // queue name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return err
	}

	// Publish message to the queue
	err = ch.Publish(
		"",         // exchange
		queue.Name, // routing key (queue name)
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Message published: %s", message)
	return nil
}

func ConsumeMessages() {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare a queue (create it if it doesn't exist)
	queue, err := ch.QueueDeclare(
		"order_queue", // queue name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Consume messages from the queue
	messages, err := ch.Consume(
		queue.Name, // queue name
		"",         // consumer tag (empty string lets RabbitMQ generate a unique tag)
		true,       // auto-ack (message acknowledgment)
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Process messages
	forever := make(chan bool)

	go func() {
		for msg := range messages {
			log.Printf("Received a message: %s", msg.Body)
			// You can handle order processing here
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}
