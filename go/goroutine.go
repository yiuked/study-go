package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Printf("CPU Number:%d", runtime.NumCPU())

	//x := true
	//c := make(chan bool) //创建一个无缓冲的bool型Channel 
	//fmt.Printf("Send...")
	//c <- x               //向一个Channel发送一个值
	//fmt.Printf("Wait...")
	//x = <- c             //从Channel c接收一个值并将其存储到x中

	x := true
	c := make(chan bool) //创建一个无缓冲的bool型Channel 
	go func() {
		log.Printf("Send...")
		c <- x               //向一个Channel发送一个值
		log.Printf("Wait...")

	}()
	close(c)
	time.Sleep(time.Second * 10)
	x = <- c             //从Channel c接收一个值并将其存储到x中


	//c := make(chan int)
	//defer close(c)
	//go func() {
	//	c<-0
	//	c<-1
	//	c<-2
	//	fmt.Println("Send data:0")
	//}()
	//time.Sleep(time.Second * 3)
	//data := <-c
	//fmt.Printf("Recv data:%d", data)
	//data = <-c
	//fmt.Printf("Recv data:%d", data)
	//data = <-c
	//fmt.Printf("Recv data:%d", data)

	//forever := make(chan int,3)
	//forever<-1
	//forever<-2
	//forever<-3
	//for  {
	//	data := <-forever
	//	fmt.Println(data)
	//	if data >=3 {
	//		break
	//	}
	//}


}
