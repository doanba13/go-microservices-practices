package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Oooopsie", http.StatusBadRequest)
		}
		log.Printf("Method %s\n Data %s\n", r.Method, d)
		fmt.Fprintf(w, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Good bye!!!")
	})

	http.ListenAndServe(":9990", nil)
}
