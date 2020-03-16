package containers

import "github.com/fwidjaya20/natsio/app-order/internal/order"

type Container struct {
	OrderService order.ServiceInterface
}

func New() Container {
	return Container{
		OrderService: order.NewOrderService(),
	}
}