package main

import (
	"fmt"
	"net/http"
)

func urlN(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "%v", req.URL.Path)
}

func main() {
	http.HandleFunc("/", urlN)
	http.ListenAndServe(":8080", nil)
}