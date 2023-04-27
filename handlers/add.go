package handlers

import (
	"net/http"

	"github.com/doanba13/data"
)

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	pd := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("%#v", pd)
	data.AddProducts(pd)

}
