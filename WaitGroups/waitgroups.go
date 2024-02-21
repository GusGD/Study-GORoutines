package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: task %s is running \n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// Thread 1: task A and B is running
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	//Thread A
	go task("A", &waitGroup)
	//Thread B
	go task("B", &waitGroup)
	//Thread anonymous
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: task %s is running \n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()
	//Exit
	waitGroup.Wait()
}
