package main

import (
	"fmt"
	"net/http"
)

// Port Number
// use const so that portNumber do not get changed by other functions
const portNumber = ":8080"

// homepage handler function
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 2)
	fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// need to make it private as it shld not be accessible outside of this package using lowercase (addValue instead of AddValue)
func addValue(x, y int) int {
	return x + y
}

func main() {
	// using the built-in package http
	// HandleFunc takes a string for the endpoint to listen to, and a handler function
	// handler func(take a http.ResponseWriter and a pointer to http.Request)
	// takes in the request and return a response
	// this handle func nvr execute as we have not start any process to listen for the process
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// use fmt.Fprintf instead of Println as Println prints to console
	// 	// takes a iowriter and a format string
	// 	n, err := fmt.Fprintf(w, "Hello, world!")

	// 	// handle if there is error
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	// })

	// refactor http handler function
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// to listen for requests use http.ListenAndServe()
	// first argument is the port to listen to
	// second argument is the handler
	// when u hit localhost on port 8080 then hello world is being printed to the browser & log entry in terminal
	_ = http.ListenAndServe(portNumber, nil)
}
