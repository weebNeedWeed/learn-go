package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

func validateFruit(fruit string) bool {
	s := []string{"apple", "pear", "banana"}
	for _, f := range s {
		if f == fruit {
			return true
		}
	}
	return false
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if t, err := template.ParseFiles("login.html"); err != nil {
			fmt.Fprint(w, "Error with html")
		} else {
			t.Execute(w, nil)
		}
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "Error when parsing form")
		return
	}

	// fmt.Println("Username:", r.Form["username"]) // r.FormValue("username")
	// .Get Only return empty value when value does not exist
	if len(r.Form.Get("username")) == 0 || len(r.Form.Get("password")) == 0 {
		fmt.Fprint(w, "username/password is not found")
		return
	}

	if _, err := strconv.Atoi(r.Form.Get("age")); err != nil {
		fmt.Fprint(w, "age is not found")
		return
	}

	if isMatched, err := regexp.MatchString(`^[a-zA-Z]$`, r.Form.Get("engname")); !isMatched || err != nil {
		fmt.Fprint(w, "engname is not found")
		return
	}

	if isMatched, err := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`,
		r.Form.Get("email")); err != nil || !isMatched {
		fmt.Fprint(w, "email is not found")
		return
	}

	if validateFruit(r.Form.Get("fruit")) {
		fmt.Fprint(w, "fruit is invalid")
		return
	}

	fmt.Println("Succeed", r.Form)
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9090", nil)
}
