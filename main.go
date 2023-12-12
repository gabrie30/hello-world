package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/healthy called")
		fmt.Fprintf(w, "healthy")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/ called")
		fmt.Fprintf(w, "hello-world")
	})

	http.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/timeout called")
		time.Sleep(10 * time.Second)
		fmt.Fprintf(w, "10 seconds")
	})

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
