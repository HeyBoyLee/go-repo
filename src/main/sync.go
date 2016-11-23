package main

import (
	"fmt"
	"sync"
)

func main	(){
	var l *sync.RWMutex
	l = new(sync.RWMutex)
	l.RLock()
	fmt.Println("runlock")
	l.RUnlock()
}