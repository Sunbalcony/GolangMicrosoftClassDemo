package main

import (
	"fmt"
	"sync"
	"time"
)

var sum = 0
var mutex sync.Mutex
var mutex2 sync.RWMutex

func main() {
	//开启100个协程让sum+10
	for i := 0; i < 10000; i++ {
		go add(10)

	}
	//防止提前退出
	time.Sleep(5 * time.Second)
	fmt.Println("和为:", sum)
	for i := 0; i < 10; i++ {
		go fmt.Println("和为:", readSum())
	}
}
func add(i int) {
	mutex.Lock()
	//defer mutex.Unlock()
	//if forget you can use defer
	sum += i
	mutex.Unlock()
}
func readSum() int {
	mutex2.RLock()
	defer mutex2.RUnlock()
	b := sum
	return b

}
