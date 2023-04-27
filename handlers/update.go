package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/doanba13/data"
	"github.com/gorilla/mux"
)

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