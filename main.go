package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format(time.RFC3339)
	fmt.Fprintf(w, "Current date & time: %s\n", now)
}

func main() {
	http.HandleFunc("/", handler)
	// Listen on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
