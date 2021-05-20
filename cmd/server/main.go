package main

import (
	"encoding/json"
	"github.com/DeOne4eg/go-event-loop/config"
	"github.com/DeOne4eg/go-event-loop/internal/calc"
	"github.com/DeOne4eg/go-event-loop/pkg/broker"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.New()

	b := broker.NewBroker(cfg, "rabbitmq")
	b.Connect()
	defer b.Close()
	b.Register("event")
	name := "event"

	go func() {
		messages := b.Consume(name)
		for msg := range messages {
			if msg.ReplyTo != "" {
				_ = b.Reject(msg.Tag, true)
				continue
			}

			log.Infof("Received a message: %v", string(msg.Body))
			var data calc.Event
			err := json.Unmarshal(msg.Body, &data)
			if err != nil {
				_ = b.Nack(msg.Tag, false, false)
				log.Fatalf("Error get json: %v", err.Error())
			}

			var result float32
			switch data.Event {
			case "calc_add":
				result = calc.Add(data.Args.A, data.Args.B)
			case "calc_sub":
				result = calc.Sub(data.Args.A, data.Args.B)
			case "calc_mut":
				result = calc.Mut(data.Args.A, data.Args.B)
			case "calc_div":
				result = calc.Divide(data.Args.A, data.Args.B)
			default:
				_ = b.Reject(msg.Tag, false)
				log.Fatalf("Unknown event: %s", data.Event)
			}

			log.Infof("Result: %.2f", result)
			r := &calc.EventResult{
				Result: result,
			}

			j, err := json.Marshal(r)
			if err != nil {
				log.Errorf("Error get JSON: %v", err.Error())
				continue
			}

			err = b.Response(msg.Tag, "event", string(j))
			if err != nil {
				_ = b.Nack(msg.Tag, false, true)
				log.Errorf("Error send response: %v", err.Error())
				continue
			}

			err = b.Ack(msg.Tag, false)
			if err != nil {
				log.Errorf("Error ack: %v", err.Error())
			}
		}
	}()

	// catch signals for quit from application
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Stopped")
}
