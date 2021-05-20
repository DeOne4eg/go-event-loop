package main

import (
	"encoding/json"
	"github.com/DeOne4eg/go-event-loop/config"
	"github.com/DeOne4eg/go-event-loop/internal/calc"
	"github.com/DeOne4eg/go-event-loop/pkg/broker"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	events = [4]string{"calc_add", "calc_sub", "calc_mut", "calc_div"}
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cfg := config.New()

	b := broker.NewBroker(cfg, "rabbitmq")
	b.Connect()
	defer b.Close()
	b.Register("event")
	name := "event"

	go func() {
		messages := b.Consume(name)
		for msg := range messages {
			if msg.ReplyTo == "" {
				_ = b.Reject(msg.Tag, true)
				continue
			}

			var r calc.EventResult
			err := json.Unmarshal(msg.Body, &r)
			if err != nil {
				_ = b.Reject(msg.Tag, false)
				log.Fatalf("Error get json: %v", err.Error())
			}

			log.Infof("[%s] Result: %.2f", msg.ReplyTo, r.Result)

			_ = b.Ack(msg.Tag, false)
		}
	}()

	for i := 0; i < 2; i++ {
		event := events[rand.Intn(len(events))]
		data := &calc.Event{
			Event: event,
			Args:  calc.EventArgs{A: rand.Intn(1000), B: rand.Intn(1000)},
		}

		d, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		err = b.Publish(name, string(d))
		if err != nil {
			log.Errorf("Error publish message: %v", err.Error())
		}
		log.Infof("Send message: %v", data)
	}

	// catch signals for quit from application
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Stopped")
}
