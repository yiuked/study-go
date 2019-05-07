package main

import "fmt"

func main() {
	global := make(map[string] interface{})

	global["name"] = "123"
	global["age"] = 123
	global["chid"] = []int{1,2,3}

	fmt.Println(global)
}

