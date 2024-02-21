package main

import "fmt"

func main() {
	channel := make(chan int)
	go publish(channel)
	reader(channel)

}
func reader(ch chan int) {
	for v := range ch {
		fmt.Printf("Received %d \n", v)
	}
}

func publish(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
