package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0
	go func() {
		for {
			atomic.AddInt64(&i, 1) //para evitar problema de concorrencia
			time.Sleep(time.Second)
			msg := Message{i, "Process 1"}
			c1 <- msg
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			msg := Message{i, "Process 2"}
			c2 <- msg
		}
	}()

	//for i := 0; i < 5; i++ {
	for { //para informações que vem de canais diferentes
		select {

		case msg := <-c1:
			fmt.Printf("ID: %d - Receiced from Process 1: %s\n", msg.Id, msg.Msg)

		case msg := <-c2:
			fmt.Printf("ID: %d - Receiced from Process 2: %s\n", msg.Id, msg.Msg)

		case <-time.After(time.Second * 3):
			println("timeout")

		}
	}
}
