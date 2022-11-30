package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	// handling the error 🔴
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() 🔴 err : %v", err)
		return
	}
	fmt.Fprintf(w, "Post request SUCCESS 🟢\n")
	// getting value from the POST request
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found 😥", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello Jeeva")
}

func main() {
	// file server to access the static html file
	fileServer := http.FileServer(http.Dir("./static"))
	// handel the static file server (default index.html)
	http.Handle("/", fileServer)
	// handel /form (function)
	http.HandleFunc("/from", formHandler)
	// handel /hello (function)
	http.HandleFunc("/hello", helloHandler)
	// server ❤️
	fmt.Println("server 🥸 :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
