package main

import "fmt"

// thread 1
func main() {
	channel := make(chan string)
	// thread 2
	go func() {
		channel <- "hello"
	}()
	//Thread 1
	msg := <-channel
	fmt.Println(msg)
}
