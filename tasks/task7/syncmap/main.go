package main

import (
	"fmt"
	"sync"
)

func main() {
	//используется потокобезопасная map из пакета sync
	var syncMap sync.Map
	//сохраняем значения. syncmap позволяет хранить пары ключ-значение разных типов
	syncMap.Store("key", "value from syncmap")
	syncMap.Store(1, 3)

	//читаем значения
	val, _ := syncMap.Load("key")
	val2, _ := syncMap.Load(1)

	fmt.Println(val)
	fmt.Println(val2)

}
