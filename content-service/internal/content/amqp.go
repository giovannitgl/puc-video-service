package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/giovannitgl/video-services/content-service/app/v1/controllers"
	"github.com/giovannitgl/video-services/content-service/internal/config"
	"github.com/giovannitgl/video-services/content-service/internal/contract"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type EventListener interface {
	Listen(eventNames ...string) (<-chan contract.Event, <-chan error, error)
}

type eventListener struct {
	connection *amqp.Connection
	queue      string
}

type eventProcessor struct {
	Listener EventListener
}

func (l *eventListener) setup() error {
	channel, err := l.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	_, err = channel.QueueDeclare(
		l.queue, // Queue name
		true,    // durable
		false,   // auto delete
		false,   // exclusive
		false,   // no wait
		nil,     // args
	)
	return err
}

func (l *eventListener) Listen(eventNames ...string) (<-chan contract.Event, <-chan error, error) {
	channel, err := l.connection.Channel()
	if err != nil {
		return nil, nil, err
	}
	defer channel.Close()
	for _, eventName := range eventNames {
		if err := channel.QueueBind(l.queue, eventName, "events", false, nil); err != nil {
			return nil, nil, err
		}
	}
	msgs, err := channel.Consume(l.queue, "", false, false, false, false, nil)
	if err != nil {
		return nil, nil, err
	}

	eventCh := make(chan contract.Event)
	errCh := make(chan error)

	go func() {
		for msg := range msgs {
			eventName, ok := msg.Headers["x-event-name"]
			if !ok {
				errCh <- errors.New("msg did not contain x-event-name header")
				msg.Nack(false, false)
				continue
			}
			eventName, ok = eventName.(string)
			if !ok {
				errCh <- fmt.Errorf("x-event-name header is not string, but %t", eventName)
				msg.Nack(false, false)
				continue
			}
			var event contract.Event
			switch eventName {
			case "video.uploaded":
				event = new(contract.VideoUploadEvent)
			default:
				errCh <- fmt.Errorf("event type %s is unknow", eventName)
				continue
			}
			err := json.Unmarshal(msg.Body, event)
			if err != nil {
				errCh <- err
				continue
			}
			eventCh <- event
		}
	}()
	return eventCh, errCh, nil
}

func (p *eventProcessor) handleEvent(event contract.Event) {
	switch e := event.(type) {
	case *contract.VideoUploadEvent:
		{
			log.Printf("event %s received", e)
			err := controllers.UpdateVideoUrl(e)
			if err != nil {
				log.Printf("failed to process event %s: %s", e, err.Error())
			}
		}
	default:
		log.Printf("unknown event: %t", e)
	}
}

func (p *eventProcessor) ProcessEvent() error {
	eventCh, errCh, err := p.Listener.Listen("video.uploaded")
	if err != nil {
		return err
	}
	for {
		select {
		case evt := <-eventCh:
			p.handleEvent(evt)
		case err = <-errCh:
			log.Printf("received error while processing msg: %s", err.Error())
		}
	}
}

func newEventListener() (EventListener, error) {
	conn, err := amqp.Dial(config.AmqpDSN())
	if err != nil {
		return nil, err
	}

	listener := &eventListener{
		connection: conn,
		queue:      config.EventQueue(),
	}
	err = listener.setup()
	if err != nil {
		return nil, err
	}
	return listener, nil
}

func NewEventProcessor() (*eventProcessor, error) {
	lis, err := newEventListener()
	if err != nil {
		return nil, err
	}

	return &eventProcessor{
		Listener: lis,
	}, nil
}
