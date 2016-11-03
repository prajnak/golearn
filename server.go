package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments in request
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello praj!") // send data to client side
}

func main() {
	http.HandleFunc("/", sayHelloName) //setup a router

	// ListenAndServe initializes a server object, calls net.Listen on a tcp
	// port to setup a tcp listener, and starts to listen on the specified
	// address and port
	err := http.ListenAndServe(":9090", nil) // set port to listen once
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
