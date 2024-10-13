package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"
)

func generateToken() string {
	currentTime := time.Now().Unix() // UTC time in second
	h := md5.New()

	// md5-ize it
	io.WriteString(h, strconv.FormatInt(currentTime, 10))
	token := h.Sum(nil)

	return fmt.Sprintf("%x", token)
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == http.MethodGet {
		t := template.Must(template.ParseFiles("login.html"))
		token := generateToken()
		t.Execute(w, struct{ Token string }{token})
		return
	}

	// Check token here
	// token := r.Form["token"]

	// fmt.Fprint(w, "")
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9090", nil)
}
