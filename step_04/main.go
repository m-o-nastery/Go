package main

import (
	"slices"
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
	val := Buf[0]
	Buf = slices.Delete(Buf, 0, 1)
	mutex.Unlock()
	return val
}