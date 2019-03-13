package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Printf("CPU Number:%d", runtime.NumCPU())

	c := make(chan int)
	go func() {
		for {
			data := <-c
			if data == 0 {
				break
			}
			fmt.Println(data)
		}
		c<-0
	}()

	for i := 10; i >= 0; i-- {
		c<-i
	}

}
