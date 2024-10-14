package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t := template.Must(template.ParseFiles("upload.html"))
		t.Execute(w, nil)
		return
	}

	r.ParseMultipartForm(32 * 1024 * 1024)

	file, fileHeader, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// fmt.Printf("%v\n", fileHeader)
	os.MkdirAll("./test", 0744)
	f, err := os.OpenFile("./test/"+fileHeader.Filename, os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	io.Copy(f, file)
}

func main() {
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":9090", nil)
}
