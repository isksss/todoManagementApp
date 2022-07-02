package router

import (
	"fmt"
	"log"
	"net/http"

	cr "github.com/isksss/todoManagementApp/crypto"

	"github.com/isksss/todoManagementApp/auth"
)

const (
	id_min int = 1000000
	id_max int = 9999999
)

func RouteIndex(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "", "_layout", "_navigate", "_footer", "index")
}

func RouteSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		generateHTML(w, "", "_layout", "_navigate", "_footer", "signup")
	} else {
		user := auth.Users{
			Id:       cr.CreateRandomInt(id_min, id_max),
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: cr.PasswdEncrypt(r.PostFormValue("password")),
		}
		fmt.Print(user)
		err := user.CreateUser()
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
}
