package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloName(w, r)
		return
	}

	http.NotFound(w, r)
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func deleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Path[len("/delete/"):]
		fmt.Println("Deleted:", id)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	// Customized Handler
	// mux := &MyMux{}
	// http.ListenAndServe(":9090", mux)

	http.HandleFunc("/delete/", deleteTaskFunc)
	http.ListenAndServe(":9090", nil)
}
