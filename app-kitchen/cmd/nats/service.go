package nats

import (
	"fmt"
	"github.com/fwidjaya20/natsio/pkg/es"
	stan "github.com/nats-io/go-nats-streaming"
)

func MakeSubscribers(conn stan.Conn) ([]stan.Subscription, error) {
	var subs []stan.Subscription
	var err error

	beverage, err := beverageSubscription(conn)
	subs = append(subs, beverage)

	food, err := foodSubscription(conn)
	subs = append(subs, food)

	return subs, err
}

func beverageSubscription(conn stan.Conn) (stan.Subscription, error) {
	var sub stan.Subscription
	var err error

	sub = es.NewSubscriber(conn, es.SubscriberData{
		Subject:    "order.created",
		QueueGroup: "order-beverage-queue",
		Durable:    "order-beverage-sub",
	}, func(msg *stan.Msg) {
		fmt.Println("FROM BEVERAGE")
		fmt.Println(msg)
	}).Subscribe()

	return sub, err
}

func foodSubscription(conn stan.Conn) (stan.Subscription, error) {
	var sub stan.Subscription
	var err error

	sub = es.NewSubscriber(conn, es.SubscriberData{
		Subject:    "order.created",
		QueueGroup: "order-food-queue",
		Durable:    "order-food-sub",
	}, func(msg *stan.Msg) {
		fmt.Println("FROM FOOD")
		fmt.Println(msg)
	}).Subscribe()

	return sub, err
}