package main

import (
	"log"
	"time"
)

func main() {
	singleReady := make(<-chan int)
	go func() {
		log.Println("Are you ok?")
		data := <-singleReady
		if data == 0 {
			log.Println("Recover")
		}
		log.Println("Are you ok?")
	}()
	log.Println("Where you?")
	time.Sleep(time.Second * 10)

	timer := time.NewTimer(time.Second)
}

