package main

import (
	"fmt"
	"sync"
	"time"
)

var lock = &sync.Mutex{}
var instance *SingleInstance

type SingleInstance struct {
}

func GetInstance() *SingleInstance {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance for the first time.")
			instance = &SingleInstance{}
		}
	}
	fmt.Println("Returning instance to the caller.")
	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		go GetInstance()
	}
	time.Sleep(time.Second * 1)
}
