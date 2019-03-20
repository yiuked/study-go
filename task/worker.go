package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type Worker struct {
	mqConn *amqp.Connection
	mqCh   map[string]*amqp.Channel
}

func (worker *Worker) initMQ() {
	var err error
	worker.mqConn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	log.Println("Connected RabbitMQ success.")
	worker.mqCh = make(map[string]*amqp.Channel)
}

func Dispatch() {
	var worker Worker
	worker.initMQ()
	defer worker.close()
	for {
		tasks := GetAll("SELECT id_task,name FROM `task`")

		fmt.Printf("Current tasks number %d\n", len(tasks))
		if nil != tasks {
			for _, task := range tasks {
				if _, ok := worker.mqCh[task["name"]]; ok {
					queueSta, _ := worker.mqCh[task["name"]].QueueInspect(task["name"])
					if queueSta.Consumers > 0 {
						continue
					}
				}

				ch, err := worker.mqConn.Channel()
				if nil != err {
					log.Printf("[error] Task[%s] create channel fail,err:%s\n", task["name"], err.Error())
					continue
				}

				worker.mqCh[task["name"]] = ch

				msg, err := addConsumer(worker.mqCh[task["name"]], task["name"])
				if nil != err {
					log.Printf("[error] Task[%s] add consumer fail,err:%s\n", task["name"], err.Error())
					worker.mqCh[task["name"]].Close()
					worker.mqCh[task["name"]] = nil
					continue
				}

				go addRecvMsg(task["name"], msg)
			}
		} else {
			log.Println("Tasks list is empty!")
		}

		time.Sleep(time.Second * 5)
	}
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

	log.Printf("Tasks[%s] is runing!\n", queueName)
	return msg, nil
}

func addRecvMsg(task string, msg <-chan amqp.Delivery) {
	log.Printf("Tasks[%s] is recving!\n", task)
	for data := range msg {
		fmt.Printf("%s\n", data.Body)
		time.Sleep(time.Second * 5)
	}
}

func (worker Worker) close() {
	worker.mqConn.Close()
}
