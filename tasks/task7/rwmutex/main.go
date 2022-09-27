package main

import (
	"fmt"
	"sync"
)

func main() {
	newSyncMap := mySyncMap{dict: make(map[string]string)}

	// запись в map
	newSyncMap.Lock()
	newSyncMap.dict["key"] = "value from map"
	newSyncMap.Unlock()

	// чтение из map
	newSyncMap.RLock()
	value := newSyncMap.dict["key"]
	newSyncMap.RUnlock()

	fmt.Println(value)
}

type mySyncMap struct {
	dict map[string]string
	sync.RWMutex
}
