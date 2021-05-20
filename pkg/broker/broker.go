package broker

import (
	"github.com/DeOne4eg/go-event-loop/config"
	"log"
)

// MQ common interface for all MQ brokers.
type MQ interface {
	Connect()
	Register(name string)
	Publish(target string, body string) error
	Response(tag uint64, target string, body string) error
	Consume(target string) <-chan Message
	Ack(tag uint64, multiple bool) error
	Nack(tag uint64, multiple bool, requeue bool) error
	Reject(tag uint64, requeue bool) error
	Close()
}

// Message is a data which received from consume.
type Message struct {
	Body        []byte
	ContentType string
	Redelivered bool
	Tag         uint64
	ReplyTo     string
}

// NewBroker create instance of Broker.
// Accept config and driver name.
func NewBroker(cfg *config.Config, driver string) MQ {
	switch driver {
	case "rabbitmq":
		return NewRabbitMQ(cfg)
	default:
		log.Fatal("Unknown broker driver")
		return nil
	}
}
