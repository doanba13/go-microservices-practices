package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/doanba13/data"
)

type KeyProduct struct{}

func (p *Products) MiddlewareValidateFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pd := &data.Product{}
		err := pd.FromJSON(r.Body)

		if err != nil {
			http.Error(w, "Cannot decode json", http.StatusBadRequest)
			return
		}
		validateErr := pd.Validate()

		if validateErr != nil {
			http.Error(w, fmt.Sprint("Validate json error: ", validateErr), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, pd)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})

}
