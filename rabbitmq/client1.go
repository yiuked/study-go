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

	// 绑定队列
	q, err := ch.QueueDeclare(
		"message_queue", // 队列名
		true,  // 是否持久
		false,  // 使用完自动删除
		false,  // 专用
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msg, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // 自动确认
		false,  // 专用
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	for data := range msg {
		fmt.Printf("%s\n", data.Body)
	}

	//forever := make(chan bool)
	//
	//go func() {
	//	for data := range msg {
	//		fmt.Printf("%s",data.Body)
	//	}
	//}()
	//
	//<-forever
}
