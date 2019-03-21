package main

import (
	"fmt"
	"reflect"
)

type Task struct {
	name string
}

func (task *Task) RepayTask(name string) {
	fmt.Println("Repay Task")
}

func (task *Task) InvestTask(name string) {
	fmt.Println("Invest Task")
}

func main()  {
	t := &Task{}
	v := reflect.ValueOf(t)

	ele := v.Elem()
	s := ele.Type()
	fmt.Println(s.NumField())

	fmt.Println(v.NumMethod())
	arg := reflect.ValueOf("ABC")
	v.MethodByName("RepayTask").Call([]reflect.Value{arg})
	v.MethodByName("InvestTask").Call([]reflect.Value{arg})
}