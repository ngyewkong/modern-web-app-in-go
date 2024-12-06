package main

import (
	"fmt"
	"net/http"
)

func main() {
	// using the built-in package http
	// HandleFunc takes a string for the endpoint to listen to, and a handler function
	// handler func(take a http.ResponseWriter and a pointer to http.Request)
	// takes in the request and return a response
	// this handle func nvr execute as we have not start any process to listen for the process
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// use fmt.Fprintf instead of Println as Println prints to console
		// takes a iowriter and a format string
		n, err := fmt.Fprintf(w, "Hello, world!")

		// handle if there is error
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	// to listen for requests use http.ListenAndServe()
	// first argument is the port to listen to
	// second argument is the handler
	// when u hit localhost on port 8080 then hello world is being printed to the browser & log entry in terminal
	_ = http.ListenAndServe(":8080", nil)
}
