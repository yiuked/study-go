package main

import (
	"fmt"
)

type Task struct {
	name string
}

// 函数名一定要大写，不然反射获取不到函数列表
func (task *Task) RepayTask(body []byte) {
	fmt.Printf("Repay Task Recv:%s\n", body)
}

// 函数名一定要大写，不然反射获取不到函数列表
func (task *Task) InvestTask(body []byte) {
	fmt.Printf("Invest Task Recv:%s\n", body)
}

//func main()  {
//	t := &Task{}
//	v := reflect.ValueOf(t)
//
//	ele := v.Elem()
//	s := ele.Type()
//	fmt.Println(s.NumField())
//
//	fmt.Println(v.NumMethod())
//	arg := reflect.ValueOf("ABC")
//	v.MethodByName("RepayTask").Call([]reflect.Value{arg})
//	v.MethodByName("InvestTask").Call([]reflect.Value{arg})
//}
