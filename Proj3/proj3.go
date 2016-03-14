package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/nu7hatch/gouuid"
)

func serve(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("proj3.html")
	if err != nil {
		log.Fatalln(err)
	}
	name := req.FormValue("Name:")
	age := req.FormValue("Age:")
	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String() + "|" + name + age,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	err = tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", serve)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}