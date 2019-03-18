package main

import (
	"log"
)

func main() {
	forever := make(chan bool)
	dataCh := make(chan string)
	var dataSingle = (<-chan string)(dataCh)
	dataCh<-"hello wrold"
	go func() {

		for d := range dataSingle {
			log.Printf("Received a message: %s", d)
		}
		log.Print("Ping ..")
	}()
	<-forever
}
