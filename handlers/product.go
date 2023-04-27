// Project title Product API.
//
// Project Products RESTFul API swagger docs.
//
//	Schemes: http
//	Version: 0.1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"log"

	"github.com/doanba13/data"
)

// Products response array
// swagger:response productsResponse
type productsResponse struct {
	// in: body
	Body []data.Product
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
