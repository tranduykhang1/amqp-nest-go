package main

import (
	"log"
	"os"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/streadway/amqp"
)

func main() {
	// Define RabbitMQ server URL.
	amqpServerURL := os.Getenv("RABBITMQ_CLIENT")
	queueName := os.Getenv("RABBITMQ_QUEUE_NAME")
	

	fmt.Println(amqpServerURL)
	fmt.Println(queueName)

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()


	
	_, err = channelRabbitMQ.QueueDeclare(
		queueName, // queue name
		true,            // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(
		logger.New(), 
	)

	app.Get("/send", func(c *fiber.Ctx) error {
		queryMsg := c.Query("msg")
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(queryMsg),
		}

		channelRabbitMQ.Publish(
			"",              // exchange
			queueName, // queue name
			false,           // mandatory
			false,           // immediate
			message,         // message to publish
		)
		return nil
	})

	log.Fatal(app.Listen(":3333"))
}