package main

import (
	"github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"encoding/base64"
)

type User struct {
	Name string
	Age string
	}

func main() {
	tpl, err := template.ParseFiles("proj4.html")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		name := req.FormValue("name")
		age := req.FormValue("age")

	person := User{
			Name: req.FormValue("name"), 
			Age: req.FormValue("age"),
		}
	
		b, err := json.Marshal(person)
		if err != nil {
			fmt.Printf("error: " , err)}
		encode := base64.StdEncoding.EncodeToString(b)


		cookie, err := req.Cookie("session-fino")
		if err != nil {
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "session-fino",
				Value: id.String() + "|" + name + age,
				Secure: true,
				HttpOnly: true,
			}
			http.SetCookie(res, cookie)

		}

		err = tpl.Execute(res, nil)
		if err != nil {
			log.Fatalln(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}