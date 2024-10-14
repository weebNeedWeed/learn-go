package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		fmt.Println(err)
		return err
	}

	bodyWriter.Close()
	contentType := bodyWriter.FormDataContentType()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func main() {
	fileName := "./coraline.png"
	targetUrl := "http://localhost:9090/upload"
	postFile(fileName, targetUrl)
}
