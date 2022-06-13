package main

import (
	"net/http"
	"log"
	// "encoding/hex"
	// "fmt"
	// "reflect"
)

func HandleIndex(w http.ResponseWriter, r *http.Request){
	generateHTML(w, "", "_layout", "_public.navigate", "_footer", "index")
}

func HandleLogin(w http.ResponseWriter, r *http.Request){
	generateHTML(w, "", "_layout", "_public.navigate", "_footer", "login")
}

func HandleLoginAuth(w http.ResponseWriter, r *http.Request){
	generateHTML(w, "", "_private.layout", "_public.navigate", "_footer", "mypage")
}

func HandleSignup(w http.ResponseWriter, r *http.Request){
	generateHTML(w, "", "_layout", "_public.navigate", "_footer", "signup")
}

func HandleSignupAuth(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	user := Users{
		name: r.FormValue("name"),
		email: r.FormValue("email"),
		password: passwdEncrypt(r.FormValue("password")),
	}

	log.Printf("NAME: %s, E-MAIL: %s, PASS: %s\n", user.name, user.email, user.password)

	if err := createUser(user); err != nil{
		panic(err)
	}
	
	generateHTML(w, "", "_layout", "_public.navigate", "_footer", "login")
}
