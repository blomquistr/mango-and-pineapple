package main

import (
	"net/http"

	m "github.com/blomquistr/mango-and-pineapple/pkg/mango"
	p "github.com/blomquistr/mango-and-pineapple/pkg/pineapple"
)

func main() {
	http.HandleFunc("/mango", m.Mango)
	http.HandleFunc("/pineapple", p.Pineapple)

	http.ListenAndServe(":8090", nil)
}
