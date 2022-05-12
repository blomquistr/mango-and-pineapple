package main

import (
	"net/http"

	"mango"
	"pineapple"
)

func main() {
	http.HandleFunc("/mango", mango.Mango)
	http.HandleFunc("/pineapple", pineapple.Pineapple)

	http.ListenAndServe(":8090", nil)
}
