package mango

import (
	"fmt"
	"net/http"
)

func Mango(w http.ResponseWriter, req *http.Request) {
	// log.Printf("[*] Got request [%s]", req)
	fmt.Fprintf(w, "Mango!")
}
