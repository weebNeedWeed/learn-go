package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == http.MethodGet {
		if t, err := template.ParseFiles("login.html"); err != nil {
			http.NotFound(w, r)
		} else {
			t.Execute(w, nil)
		}
		return
	}

	fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
	fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))

	template.HTMLEscape(w, []byte(r.Form.Get("username")))

	fmt.Fprintln(w)

	// Print html
	t, _ := template.New("t").Parse(`{{define "T"}}Hello {{.}}!{{end}}`)
	// without template.HTML, it will print the escaped string
	t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9090", nil)
}
