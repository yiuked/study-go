package main

import (
	"fmt"
	"reflect"
)

//var eventByName = make(map[string][]func(interface{}))
//
//func funcName1(){
//	fmt.Print("funcName1")
//}
//func funcName2(){
//	fmt.Print("funcName2")
//}
//func funcName3(){
//	fmt.Print("funcName3")
//}
//
//func RegisterEvent(name string, callback func(interface{})) {
//	// 通过名字查找事件列表
//	list := eventByName[name]
//	// 在列表切片中添加函数
//	list = append(list, callback)
//	// 将修改的事件列表切片保存回去
//	eventByName[name] = list
//}
//
//// 调用事件
//func CallEvent(name string, param interface{}) {
//	// 通过名字找到事件列表
//	list := eventByName[name]
//	// 遍历这个事件的所有回调
//	for _, callback := range list {
//		// 传入参数调用回调
//		callback(param)
//	}
//}
//func main() {
//	RegisterEvent("funcName1", func(i interface{}) {
//
//	})
//}

func main()  {
	var task = func() {
		fmt.Println("...")
	}
	to := reflect.TypeOf(task)
	fmt.Println(to.Name())
	fmt.Println(to.String())
	vo := reflect.ValueOf(task)
	vo.ca
	fmt.Println(vo)
}
