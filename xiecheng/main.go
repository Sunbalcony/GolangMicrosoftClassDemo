package main

import (
	"fmt"
	"time"
)

func downloadFile(chanName string) string {
	time.Sleep(time.Second * 2)
	return chanName + ":filePath"

}
func main() {
	//声明三个存放结果的channel
	firstCh := make(chan string)
	secondCh := make(chan string)
	thirdCh := make(chan string)
	go func() {
		firstCh <- downloadFile("firstCh")
	}()
	go func() {
		firstCh <- downloadFile("secondCh")
	}()
	go func() {
		firstCh <- downloadFile("thirdCh")
	}()
	//开始select多路复用，哪个channel能最先获取到结果,就代表哪个最好
	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-thirdCh:
		fmt.Println(filePath)

	}

}
