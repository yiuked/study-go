package main

import "fmt"

func main() {

	name,age,sex := foo(123)
	fmt.Printf("%s:%d:%d",name,age,sex)

	name1,_,sex3 := foo(123, "李四")
	fmt.Printf("%s:%d",name1,sex3)

}

func foo(user_id int, args...interface{}) (name string, age int, sex int)  {
	return "张三",25, 1
}