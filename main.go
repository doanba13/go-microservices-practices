package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/doanba13/handlers"
)

func main() {
	log := log.New(os.Stdout, "product-api: ", log.LstdFlags)
	helloHandler := handlers.NewHello(log)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)

	productHandler := handlers.NewProducts(log)
	serveMux.Handle("/product/", productHandler)

	s := &http.Server{
		Addr:         ":9990",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)

	sig := <-signalChan
	log.Println("Recieved terminal, grateful shutdown in 30s", sig)

	timeOutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	s.Shutdown(timeOutContext)
}
