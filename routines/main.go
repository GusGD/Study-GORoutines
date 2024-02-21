package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: task %s is running \n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1: task A and B is running
func main() {
	//Thread A
	go task("A")
	//Thread B
	go task("B")
	//Thread anonymous
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: task %s is running \n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()
	//Exit
	time.Sleep(15 * time.Second)
}
