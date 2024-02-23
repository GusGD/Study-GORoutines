package main

func main() {

	channel := make(chan int, 2)
	channel <- 1
	channel <- 2

	print(<-channel)
	print(<-channel)
}
