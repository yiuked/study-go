package main

import (
	"fmt"
	"time"
)

func main() {

	var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	t1,_ := time.ParseInLocation("2006-01-02 15:04:05", "2019-04-26 11:38:00", cstSh)
	fmt.Println(t1)
	t2 := time.Now()
	fmt.Println(t2)
	fmt.Println(t2.Sub(t1).Seconds())
	if t2.Sub(t1).Seconds() > 100 {
		fmt.Println("OK")
	} else {
		fmt.Println("Err")
	}
}