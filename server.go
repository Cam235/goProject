package main

//import needed packages
import (
	"fmt"
	"log"
	"net/http"
)

// main function to start the server
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// helloHandler function to handle the request
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

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error parsing form: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successfully received.")
	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Name: %s\nEmail: %s\n", name, email)
}
