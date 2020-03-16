package es

import (
	"bytes"
	"encoding/json"
	stan "github.com/nats-io/go-nats-streaming"
)

var globalPublisher *Publisher

type Publisher struct {
	conn stan.Conn
}

type StoreData struct {
	Channel     string
	Domain      string
	Subject     string
	EventSource string
	Data        interface{}
}

func NewPublisher(conn stan.Conn) *Publisher {
	return &Publisher{
		conn: conn,
	}
}

func (p *Publisher) Store(info StoreData) error {
	var t bytes.Buffer
	var result = make(map[string]interface{}, 0)

	result["domain"] = info.Domain
	result["subject"] = info.Subject
	result["event_source"] = info.EventSource
	result["data"] = info.Data

	resultByte, err := json.Marshal(result)

	if nil != err {
		print("__METHOD__=Store", "Transport=NATS", "error", err)
		panic(err)
	}

	t.Write(resultByte)

	err = p.conn.Publish(info.Subject, t.Bytes())

	if nil == err {
		print("__METHOD__=Store", "Transport=NATS", "__PUBLISHED_MESSAGE_ON__", info.Channel, "__WITH SUBJECT__", info.Subject, )
	} else {
		print("__METHOD__=Store", "Transport=NATS", "error", err)
	}

	return err
}

func SetGlobalPublisher(publisher *Publisher) {
	globalPublisher = publisher
}

func GetGlobalPublisher() *Publisher {
	return globalPublisher
}