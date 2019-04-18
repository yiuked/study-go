package main

import (
	"plugin"
)

func main() {
	p, err := plugin.Open("test.so")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("Hello")
	if err != nil {
		panic(err)
	}
	f.(func())()
}