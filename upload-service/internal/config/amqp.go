package config

import (
	"context"
	"encoding/json"
	"github.com/giovannitgl/video-services/content-service/internal/contract"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type EventPublisher interface {
	Publish(event contract.Event) error
}

type eventPublisher struct {
	connection *amqp.Connection
}

func (p *eventPublisher) setup() error {
	channel, err := p.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	return channel.ExchangeDeclare(
		"events",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}

func (p *eventPublisher) Publish(event contract.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	ch, err := p.connection.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()
	msg := amqp.Publishing{
		Headers:     amqp.Table{"x-event-name": event.EventName()},
		Body:        data,
		ContentType: "application/json",
	}
	ctx := context.Background()
	return ch.PublishWithContext(
		ctx,
		"events",
		event.EventName(),
		false,
		false,
		msg,
	)
}

func NewEventPublisher() (EventPublisher, error) {
	conn, err := amqp.Dial(AmqpDSN())
	if err != nil {
		return nil, err
	}

	publisher := &eventPublisher{
		connection: conn,
	}
	err = publisher.setup()
	if err != nil {
		return nil, err
	}
	return publisher, nil
}

var Pub EventPublisher

func SetupPublisher() {
	pub, err := NewEventPublisher()
	if err != nil {
		log.Panic(err)
	}
	Pub = pub
}
