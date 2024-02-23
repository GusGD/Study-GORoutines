package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d \n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	QtdWorkers := 100000

	for i := 0; i < QtdWorkers; i++ {
		go worker(i, data)
	}
	for i := 0; i < 1000000; i++ {
		data <- i
	}
}
