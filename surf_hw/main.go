package main

import (
	"net/http"
	"html/template"
	"log"
)

func makesite(res http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(res, nil)
}
func main() {

	http.HandleFunc("/", makesite)
	http.Handle("/pics/", http.StripPrefix("/pics", http.FileServer(http.Dir("./pics"))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))

	http.ListenAndServe(":8040", nil)
}