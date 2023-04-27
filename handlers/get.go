package handlers

import (
	"net/http"

	"github.com/doanba13/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 	 200: productsResponse

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList := data.GetListProduct()

	err := productList.ToJSON(w)

	if err != nil {
		http.Error(w, "Can't marshal production list!", http.StatusInternalServerError)
	}
}
