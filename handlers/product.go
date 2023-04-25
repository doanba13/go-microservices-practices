package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

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

	if r.Method == http.MethodPost {
		p.addProducts(w, r)
		return
	}

	if r.Method == http.MethodPut {
		p.updateProduct(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) updateProduct(w http.ResponseWriter, r *http.Request) {
	reg := regexp.MustCompile(`/([0-9]+)`)
	matchesReg := reg.FindAllStringSubmatch(r.URL.Path, -1)

	if len(matchesReg) != 1 {
		http.Error(w, "Invalid url!", http.StatusBadRequest)
		return
	}

	if len(matchesReg[0]) != 2 {
		http.Error(w, "More than two id in path!", http.StatusBadRequest)
		return
	}

	idStr := matchesReg[0][1]

	id, err := strconv.Atoi(idStr)

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

	newP := data.CreateProduct(w, r.Body)

	oldPd.Name = newP.Name
	oldPd.Description = newP.Description
	oldPd.Price = newP.Price
	oldPd.UpdatedOn = time.Now().UTC().String()
}

func (p *Products) addProducts(w http.ResponseWriter, r *http.Request) {
	pd := data.CreateProduct(w, r.Body)
	if pd != nil {
		data.AddProducts(pd)
		p.l.Printf("%#v", pd)
	}
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	productList := data.GetListProduct()

	err := productList.ToJSON(w)

	if err != nil {
		http.Error(w, "Can't marshal production list!", http.StatusInternalServerError)
	}
}
