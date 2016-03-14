package main

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", Cook)
	http.ListenAndServe(":8080", nil)
}

func Cook(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("session-fino")

	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String(),
			 Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)

}