package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

var visitorNumber int
var mux sync.Mutex

func main() {

	http.HandleFunc("/", HelloHandler)

	fmt.Println("Listening on localhost:80")
	http.ListenAndServe(":80", nil)
}

// HelloHandler Handle HTTP request
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("unable to get hostname")
	}
	version := os.Getenv("VERSION")

	mux.Lock()
	visitorNumber++
	fmt.Fprintf(w, "Hello from Go on %s, server version is: %s, you are visitor: %d\n", hostname, version, visitorNumber)
	mux.Unlock()
}
