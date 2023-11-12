package main

//import needed packages
import (
	"fmt"
	"log"
	"net/http"
)

//main function to start the server
func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

//helloHandler function to handle the request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Invalid URL.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello World!")
}
