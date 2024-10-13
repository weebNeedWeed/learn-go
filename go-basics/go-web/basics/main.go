package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error %v", err)
		return
	}

	fmt.Println(r.Form)
	fmt.Println("Scheme", r.URL.Scheme)
	fmt.Println("Path", r.URL.Path)

	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("Key:", k)
		fmt.Println("Val:", strings.Join(v, "-"))
	}

	fmt.Fprintf(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", sayHelloName)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
