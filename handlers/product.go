package handlers

import (
	"log"
	"net/http"

	"github.com/doanba13/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	productList := data.GetListProduct()

	err := productList.ToJSON(w)

	if err != nil {
		http.Error(w, "Can't marshal production list!", http.StatusInternalServerError)
	}
}
