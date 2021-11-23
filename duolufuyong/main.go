package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing"

}
func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done Replicating"

}
func third(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "Done Third"

}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	go process(ch1)
	go replicate(ch2)
	go third(ch3)
	for i := 0; i < 3; i++ {
		select {
		case a1 := <-ch1:
			fmt.Println(a1)
		case a2 := <-ch2:
			fmt.Println(a2)
		case a3 := <-ch3:
			fmt.Println(a3)


		}
	}
}
