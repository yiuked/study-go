package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Task struct {
	Queue string
	Name string
}

func main() {
	configFile, err := ioutil.ReadFile("resource/config.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
	}

	var tasks *[]Task
	err = yaml.Unmarshal(configFile, &tasks)

	for task := range tasks {
		fmt.Println(task.Queue)
	}
}
