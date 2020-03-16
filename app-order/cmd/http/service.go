package http

import (
	"encoding/json"
	"github.com/fwidjaya20/natsio/app-order/cmd/containers"
	"github.com/fwidjaya20/natsio/app-order/internal/order/data/requests"
	"github.com/go-chi/chi"
	"net/http"
)

func MakeHandler(router *chi.Mux, services containers.Container) *chi.Mux {
	router.Post("/order", func(w http.ResponseWriter, r *http.Request) {
		service := services.OrderService

		var err error
		var payload requests.OrderRequest

		defer r.Body.Close()
		err = json.NewDecoder(r.Body).Decode(&payload)

		if nil != err {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			_ = json.NewEncoder(w).Encode(err.Error())
			return
		}

		ok, err := service.Order(r.Context(), payload)

		if nil != err {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			_ = json.NewEncoder(w).Encode(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(204)
		_ = json.NewEncoder(w).Encode(ok)
	})

	return router
}