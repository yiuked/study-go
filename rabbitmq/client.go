package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"runtime"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	fmt.Printf("Num Goroutine: %d\n", runtime.NumGoroutine())

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()
	fmt.Printf("Num Goroutine: %d\n", runtime.NumGoroutine())

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	fmt.Printf("Num Goroutine: %d\n", runtime.NumGoroutine())

	msgs := make(map[string]<-chan amqp.Delivery)
	for i := 0; i < 5; i++ {
		qName := fmt.Sprintf("task_%d", i)
		msg, _ := addConsumer(ch, qName)
		if nil != msg {
			msgs[qName] = msg
		}
	}
	fmt.Printf("Num Goroutine: %d\n", runtime.NumGoroutine())


	for taskName, msg := range msgs {
		taskQueue, _ := ch.QueueInspect(taskName)
		fmt.Printf("Queue:%s,Consumers:%d\n", taskName, taskQueue.Consumers)
		go addRecvMsg(taskName, msg)
	}
	fmt.Printf("Num Goroutine: %d\n", runtime.NumGoroutine())



}

func addConsumer(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if nil != err {
		return nil, err
	}

	msg, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if nil != err {
		return nil, err
	}

	return msg, nil
}

func addRecvMsg(task string, msg <-chan amqp.Delivery)  {
	fmt.Printf("%s", task)
	for data := range msg {
		fmt.Printf("%s\n", data.Body)
		time.Sleep(time.Second * 5)
	}
}


