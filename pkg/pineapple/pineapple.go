package pineapple

import (
	"fmt"
	"net/http"
)

func Pineapple(w http.ResponseWriter, req *http.Request) {
	// log.Printf("[*] Got request [%s]", req)
	fmt.Fprintf(w, "Pineapple!")
}
