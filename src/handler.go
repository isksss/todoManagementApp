package main

import (
	"log"
	"net/http"
	// "encoding/hex"
	// "fmt"
	// "reflect"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "", "_layout", "_public.navigate", "_footer", "index")
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "", "_layout", "_public.navigate", "_footer", "login")
}

func HandleLoginAuth(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "", "_private.layout", "_public.navigate", "_footer", "mypage")
}

func HandleSignup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "", "_layout", "_public.navigate", "_footer", "signup")
}

func HandleSignupAuth(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	user := Users{
		name:     r.PostFormValue("name"),
		email:    r.PostFormValue("email"),
		password: passwdEncrypt(r.PostFormValue("password")),
	}

	log.Printf("NAME: %s, E-MAIL: %s, PASS: %s\n", user.name, user.email, user.password)

	if err := user.createUser(); err != nil {
		panic(err)
	}

	generateHTML(w, "", "_layout", "_public.navigate", "_footer", "login")
}

/***** ***** ***** ***** *****
TestCode:
***** ***** ***** ***** *****/
func HandleTest(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// fmt.Println(r.GetBody)
	if r.Method != "GET" {
		log.Printf("Method is POST")
		test := struct {
			METHOD string
			HOST   string
			NAME   string
			EMAIL  string
			PASSWD string
			// WORLD  string
			// NAME2  string
		}{
			METHOD: r.Method,
			HOST:   r.Host,
			NAME:   r.PostFormValue("name"),
			EMAIL:  r.PostFormValue("email"),
			PASSWD: r.PostFormValue("password"),
			// NAME:   r.PostForm["name"][0],
			// EMAIL:  r.PostForm["email"][0],
			// PASSWD: r.PostForm["password"][0],
			// WORLD: r.Form["world"][0],
			// NAME2: r.Form["name2"][0],
		}
		log.Printf("KEY-MAP: ")
		generateHTML(w, test, "__test.layout", "_public.navigate", "_footer", "__test")
	} else {
		log.Printf("Method is GET")
		generateHTML(w, "", "__test.layout", "_public.navigate", "_footer", "__test.input")
	}
}
