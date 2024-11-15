package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/fprint", func(w http.ResponseWriter, r *http.Request) {
		customerFile, err := os.Open("./customers.csv")
		if err != nil {
			log.Fatal(err)
		}

		defer customerFile.Close()

		// data, err := io.ReadAll(customerFile)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Fprint(w, string(data))
		
		// using io.Copy() instead of io.ReadAll() and fmt.Fprint()
		io.Copy(w, customerFile)
	})

	s := http.Server{
		Addr: ":3000",
	}

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
