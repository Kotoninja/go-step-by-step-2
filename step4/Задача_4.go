package main

import (
	"sync"
)

var (
	Buf   []int
	mutex sync.Mutex
)

func Write(num int) {
	mutex.Lock()
	Buf = append(Buf, num)
	mutex.Unlock()
}

func Consume() int {
	mutex.Lock()
	var value int
	if len(Buf) > 0 {
		value = Buf[0]
		Buf = Buf[1:]
	}
	mutex.Unlock()
	return value
}
