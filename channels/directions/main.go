package main

import "fmt"

func receives(name string, hello chan<- string) {
	hello <- name
}

func sends(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go receives("hello", hello)
	sends(hello)
}
