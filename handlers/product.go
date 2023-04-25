package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/doanba13/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "id not valid!", http.StatusBadRequest)
		return
	}

	p.l.Println("Product id", id)
	oldPd := data.GetProductById(id)

	if oldPd == nil {
		http.Error(w, "Not Found Product!", http.StatusNotFound)
		return
	}

	newP := r.Context().Value(KeyProduct{}).(*data.Product)

	oldPd.Name = newP.Name
	oldPd.Description = newP.Description
	oldPd.Price = newP.Price
	oldPd.UpdatedOn = time.Now().UTC().String()
}

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	pd := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("%#v", pd)
	data.AddProducts(pd)

}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList := data.GetListProduct()

	err := productList.ToJSON(w)

	if err != nil {
		http.Error(w, "Can't marshal production list!", http.StatusInternalServerError)
	}
}

type KeyProduct struct{}

func (p *Products) MiddlewareValidateFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pd := &data.Product{}
		err := pd.FromJSON(r.Body)

		if err != nil {
			http.Error(w, "Cannot decode json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, pd)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})

}
