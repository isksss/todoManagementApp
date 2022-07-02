package handler

import "net/http"

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "", "_layout", "_public.navigate", "_footer", "index")
}
