package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", myHandler("Customer service"))

	var handlerFunc http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.String())
	}

	http.HandleFunc("/url/", handlerFunc)

	// http.HandleFunc("/service/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Customer service")
	// })

	s := http.Server{
		Addr: ":3000",
	}

	// run the following command
	// go run 'C:\Program Files\Go\src\crypto\tls\generate_cert.go' --host localhost
	// to generate selfsigned cert.pem and key.pem for localhost
	go func() {
		log.Fatal(s.ListenAndServeTLS("./cert.pem", "./key.pem"))
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}

type myHandler string

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "energetic gophers")
	fmt.Fprintln(w, string(mh))
	fmt.Fprintln(w, r.Header)
}
