package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Worker struct {
	mqConn *amqp.Connection
	mqCh   *amqp.Channel
}

func (worker *Worker) initMQ() {
	var err error
	worker.mqConn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	log.Println("Connected RabbitMQ success.")
	worker.mqCh, err = worker.mqConn.Channel()
	failOnError(err, "Failed to open a channel")
	log.Println("Opened RabbitMQ channel.")
}

func Dispatch() {
	var worker Worker
	worker.initMQ()
	defer worker.close()
	go func() {
		tasks := GetAll("SELECT id_task,name FROM `task`")
		if nil != tasks {
			for _, task := range tasks {
				q, err := worker.mqCh.QueueDeclare(
					task["name"], // name
					false,        // durable
					false,        // delete when unused
					false,        // exclusive
					false,        // no-wait
					nil,          // arguments
				)
				failOnError(err, "Failed to declare a queue")

				data, err := worker.mqCh.Consume(
					q.Name, // queue
					"",     // consumer
					true,   // auto-ack
					false,  // exclusive
					false,  // no-local
					false,  // no-wait
					nil,    // args
				)
				failOnError(err, "Failed to register a consumer")
				go func() {
					for msg := range data {
						fmt.Printf("[%s] :%s\n", q.Name, msg.Body)
					}
				}()
			}
		} else {
			log.Println("Tasks list is empty!")
		}
	}()
}

func (worker Worker) close() {
	worker.mqConn.Close()
	worker.mqCh.Close()
}
