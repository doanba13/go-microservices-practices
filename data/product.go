package data

import (
	"encoding/json"
	"io"
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

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func AddProducts(p *Product) {
	p.ID = getNextSlideId()
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()
	p.DeletedOn = time.Now().UTC().String()
	productList = append(productList, p)
}

func getNextSlideId() int {
	pl := GetListProduct()
	lastPd := pl[len(pl)-1]
	return lastPd.ID + 1
}

func GetListProduct() Products {
	return productList
}

func GetProductById(id int) *Product {
	pl := GetListProduct()
	for i := range pl {
		if pl[i].ID == id {
			return pl[i]
		}
	}
	return nil
}

func CreateProduct(w http.ResponseWriter, r io.Reader) *Product {
	pd := &Product{}
	err := pd.FromJSON(r)

	if err != nil {
		http.Error(w, "Cannot decode json", http.StatusBadRequest)
		return nil
	}

	return pd
}

var productList = Products{
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
