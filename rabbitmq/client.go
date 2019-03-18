package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	var q amqp.Queue
	var msgs []<-chan amqp.Delivery
	for i := 0; i < 5; i++ {
		q, err = ch.QueueDeclare(
			fmt.Sprintf("task_%d", i), // name
			false,                     // durable
			false,                     // delete when unused
			false,                     // exclusive
			false,                     // no-wait
			nil,                       // arguments
		)
		failOnError(err, "Failed to declare a queue")

		msg, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		msgs = append(msgs, msg)
		failOnError(err, "Failed to register a consumer")
	}

	forever := make(chan bool)
	for k, msg := range msgs {
		go func() {
			for data := range msg {
				fmt.Printf("%d", k)
				fmt.Printf("%s",data.Body)
			}
		}()
	}
	<-forever
}
