package main

import (
	"fmt"
	"net/http"
)

func formValue(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "%v", req.FormValue("n"))
}

func main() {
	http.HandleFunc("/", formValue)
	http.ListenAndServe(":8080", nil)
}