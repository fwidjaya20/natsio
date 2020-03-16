package es

import (
	stan "github.com/nats-io/go-nats-streaming"
)

type Subscriber struct {
	conn stan.Conn
	data       SubscriberData
	handler    stan.MsgHandler
}

type SubscriberData struct {
	Subject    string
	QueueGroup string
	Durable    string
}

func NewSubscriber(
	conn stan.Conn,
	data SubscriberData,
	handler stan.MsgHandler,
) *Subscriber {
	return &Subscriber{
		conn:    conn,
		data:    data,
		handler: handler,
	}
}

func (s *Subscriber) Subscribe() stan.Subscription {
	var opts []stan.SubscriptionOption

	opts = append(opts, stan.DurableName(s.data.Durable))
	opts = append(opts, stan.DeliverAllAvailable())

	sub, err := s.conn.QueueSubscribe(s.data.Subject, s.data.QueueGroup, s.handler, opts...)

	if nil != err {
		print("_METHOD_=Subscribe", "Transport=NATS", "error", err)
		return nil
	}

	return sub
}