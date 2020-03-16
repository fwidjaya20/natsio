package order

import (
	"context"
	"fmt"
	"github.com/fwidjaya20/natsio/app-order/internal/order/data/requests"
	"github.com/fwidjaya20/natsio/pkg/es"
)

type ServiceInterface interface {
	Order(context.Context, requests.OrderRequest) (bool, error)
}

type service struct {
	actor string
}

func NewOrderService() ServiceInterface {
	return &service{
		actor: "ORDER",
	}
}

func (s *service) Order(ctx context.Context, request requests.OrderRequest) (bool, error) {
	var result bool
	var err error

	fmt.Printf("PROCESSING ORDER [%s]...\n", request.Code)
	fmt.Println("DONE")

	err = s.publishOrderCreated(ctx, request)

	if nil != err {
		fmt.Println("ERROR_Publishing_Order", "Error", err)
	}

	return result, err
}

func (s *service) publishOrderCreated(ctx context.Context, data interface{}) error {
	publisher := es.GetGlobalPublisher()
	
	return publisher.Store(es.StoreData{
		Channel:     "order",
		Domain:      "order",
		Subject:     "order.created",
		EventSource: "order.order",
		Data:        data,
	})
}