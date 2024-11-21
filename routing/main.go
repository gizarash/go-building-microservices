package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Product struct {
	ID         int
	Name       string
	USDPerUnit float64
	Unit       string
}

func main() {

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(products)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-type", "application/json")
		w.Write(data)
	})

	// http://localhost:4000/products/1
	pattern := regexp.MustCompile(`^\/products\/(\d+?)$`)
	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		matches := pattern.FindStringSubmatch(r.URL.Path) // [ "/products/1" "1" ]
		if len(matches) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, p := range products {
			if p.ID == id {
				data, err := json.Marshal(p)
				if err != nil {
					log.Print(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Header().Add("Content-type", "application/json")
				w.Write(data)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	})

	s := http.Server{
		Addr: ":4000",
	}

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")

}
