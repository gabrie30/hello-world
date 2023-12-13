package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	// Handle "/timeout" endpoint
	// Example usage: /timeout?sleep=5
	http.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/timeout called")
		// Get the sleep time from the query parameters
		sleepTimeStr := r.URL.Query().Get("sleep")
		if sleepTimeStr != "" {
			// Convert the sleep time to an integer
			sleepTime, err := strconv.Atoi(sleepTimeStr)
			if err != nil {
				fmt.Println("Invalid sleep time provided")
				fmt.Fprintf(w, "Invalid sleep time provided")
			} else {
				// Sleep for the specified amount of time
				fmt.Printf("Sleeping for %d seconds\n", sleepTime)
				time.Sleep(time.Duration(sleepTime) * time.Second)
				fmt.Fprintf(w, "Slept for %d seconds", sleepTime)
			}
		} else {
			// If no sleep time is provided, default to 10 seconds
			fmt.Println("No sleep time provided, defaulting to 10 seconds")
			time.Sleep(10 * time.Second)
			fmt.Fprintf(w, "Slept for 10 seconds")
		}
	})

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
