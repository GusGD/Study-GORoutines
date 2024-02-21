package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)
	go publish(channel)
	go reader(channel, &waitGroup)
	waitGroup.Wait()
}
func reader(ch chan int, waitGroup *sync.WaitGroup) {
	for v := range ch {
		fmt.Printf("Received %d \n", v)
		waitGroup.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
