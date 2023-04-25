package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/doanba13/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log := log.New(os.Stdout, "product-api: ", log.LstdFlags)

	newMux := mux.NewRouter()
	getMux := newMux.Methods(http.MethodGet).Subrouter()
	putMux := newMux.Methods(http.MethodPut).Subrouter()
	postMux := newMux.Methods(http.MethodPost).Subrouter()

	productHandler := handlers.NewProducts(log)

	getMux.HandleFunc("/product", productHandler.GetProducts)

	putMux.HandleFunc("/product/{id:[0-9]+}", productHandler.UpdateProduct)
	putMux.Use(productHandler.MiddlewareValidateFunc)

	postMux.HandleFunc("/product", productHandler.AddProducts)
	postMux.Use(productHandler.MiddlewareValidateFunc)

	s := &http.Server{
		Addr:         ":9990",
		Handler:      newMux,
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
