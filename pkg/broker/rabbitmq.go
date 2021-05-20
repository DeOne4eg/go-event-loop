package broker

import (
	"fmt"
	"github.com/DeOne4eg/go-event-loop/config"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"strconv"
)

// RabbitMQ contains vars for control broker.
type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	cfg  *config.Config
}

// NewRabbitMQ create instance of RabbitMQ.
func NewRabbitMQ(cfg *config.Config) *RabbitMQ {
	return &RabbitMQ{cfg: cfg}
}

// Connect to RabbitMQ.
func (r *RabbitMQ) Connect() {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", r.cfg.Broker.Username, r.cfg.Broker.Password, r.cfg.Broker.Host, r.cfg.Broker.Port))
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	r.ch = ch
	r.conn = conn
}

// Register create queue.
func (r *RabbitMQ) Register(name string) {
	_, err := r.ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
}

// Consume messages from broker.
func (r RabbitMQ) Consume(target string) <-chan Message {
	messages, err := r.ch.Consume(
		target, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	ch := make(chan Message)
	go func() {
		for d := range messages {
			ch <- Message{
				Body:        d.Body,
				ContentType: d.ContentType,
				Redelivered: d.Redelivered,
				Tag:         d.DeliveryTag,
				ReplyTo:     d.ReplyTo,
			}
		}
	}()

	return ch
}

// Publish sending message to rabbit.
func (r *RabbitMQ) Publish(target string, body string) error {
	err := r.ch.Publish(
		"",
		target,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	return err
}

// Response send answer to message.
func (r *RabbitMQ) Response(tag uint64, target string, body string) error {
	err := r.ch.Publish(
		"",
		target,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			ReplyTo:     strconv.FormatUint(tag, 10),
			Body:        []byte(body),
		},
	)

	return err
}

// Ack accept message by tag.
func (r *RabbitMQ) Ack(tag uint64, multiple bool) error {
	return r.ch.Ack(tag, multiple)
}

// Nack decline message by tag with possible to requeue.
func (r *RabbitMQ) Nack(tag uint64, multiple bool, requeue bool) error {
	return r.ch.Nack(tag, multiple, requeue)
}

// Reject message by tag with possible to requeue.
func (r *RabbitMQ) Reject(tag uint64, requeue bool) error {
	return r.ch.Reject(tag, requeue)
}

// Close connection.
func (r *RabbitMQ) Close() {
	_ = r.ch.Close()
	_ = r.conn.Close()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
