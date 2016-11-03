package main

import (
	"fmt"
	"net/http"
)

// GO uses goroutines for every job initiated by Conn in order to achieve high
// concurrency and performance.

// default router structure
//
// type ServeMux struct {
// 	mu sync.RWMutex
// 	m  map[string]muxEntry
// }

// type muxEntry struct {
// 	explicit bool
// 	h        Handler
// }

// type handler interface {
// 	ServeHTTP(ResponseWriter, *Request) //routing implement
// }

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello route!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
