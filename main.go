package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 80
	}

	var count int64

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newCount := atomic.AddInt64(&count, 1)

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "There have been %d invocations!", newCount)
	})

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Print("Listening on " + addr)
	http.ListenAndServe(addr, handler)
}
