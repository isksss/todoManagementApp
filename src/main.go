package main 

import (
	"net/http"
	"log"
)

func main(){

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public/"))))

	port := "8080" //ポート番号
	server := http.Server{
		Addr: ":"+port,
	}

	// ハンドラー
	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/login", HandleLogin)
	http.HandleFunc("/login/auth", HandleLoginAuth)
	http.HandleFunc("/signup", HandleSignup)
	http.HandleFunc("/signup/auth", HandleSignupAuth)
	http.HandleFunc("/test", HandleTest)

	log.Printf("Server listening on port %s", port)
	server.ListenAndServe()
}