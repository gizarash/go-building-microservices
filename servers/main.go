package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Customer service")
	})

	// run the following command
	// go run 'C:\Program Files\Go\src\crypto\tls\generate_cert.go' --host localhost
	// to generate selfsigned cert.pem and key.pem for localhost
	log.Fatal(http.ListenAndServeTLS(":3000", "./cert.pem", "./key.pem", nil))
}
