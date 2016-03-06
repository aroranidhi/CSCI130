package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseFiles("name.html")
}

func name(res http.ResponseWriter, req *http.Request) {
	tpl.Execute(res, nil)
	fmt.Fprintf(res, "%v", req.FormValue("n"))
}

func main() {
	http.HandleFunc("/", name)
	http.ListenAndServe(":8080", nil)
}