package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseFiles("fileupload.html")
}

func urlName(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		file, _, err := req.FormFile("n")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer file.Close()

		src := io.LimitReader(file, 400)

		dst, err := os.Create(filepath.Join(".", "file.txt"))
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()

		io.Copy(dst, src)
	}
	tpl.Execute(res, nil)
	file, _ := ioutil.ReadFile("file.txt")
	fmt.Fprint(res, strings.Split(string(file), "\n"))
}

func main() {
	http.HandleFunc("/", urlName)
	http.ListenAndServe(":8080", nil)
}