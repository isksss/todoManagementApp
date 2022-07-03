package main

import (
	"log"
	"net/http"

	r "github.com/isksss/todoManagementApp/router"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public/"))))

	port := "8080" //ポート番号
	server := http.Server{
		Addr: ":" + port,
	}

	// ハンドラー
	http.HandleFunc("/", r.RouteIndex)
	http.HandleFunc("/signup", r.RouteSignUp)
	http.HandleFunc("/login", r.RouteLogin)

	log.Printf("Server listening on port %s", port)
	server.ListenAndServe()
}
