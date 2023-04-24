package data

import (
	"encoding/json"
	"net/http"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(rw http.ResponseWriter) error {
	e := json.NewEncoder(rw)

	return e.Encode(p)
}

func GetListProduct() Products {
	return productList
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Milk coffee",
		Description: "Coffee and milk, i don't know",
		Price:       4.32,
		SKU:         "GONGCHA",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso coffee",
		Description: "Coffee and more milk, i don't care",
		Price:       4.32,
		SKU:         "GONGCHA",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}
