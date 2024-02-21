package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

// ab -n 10000 -c 100 http://localhost:3000/
var number uint64 = 0

func main() {
	//m := sync.Mutex{} --- MUTEX
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		//number++
		atomic.AddUint64(&number, 1)
		//m.Unlock()
		time.Sleep(100 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}
