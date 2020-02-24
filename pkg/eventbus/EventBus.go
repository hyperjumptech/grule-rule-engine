package eventbus

import (
	"fmt"
	"github.com/imkira/go-observer"
	"sync"
)

var (
	// DefaultBrooker is a system wide Broker. To easily help developer use the event bus straight away.
	DefaultBrooker *Broker
)

func init() {
	DefaultBrooker = NewBroker()
}

// Subscriber the subscriber of a broker. This subscriber already asigned with topic it should subscribe.
type Subscriber struct {
	// Topic that subscribed.
	Topic string
	// Handle function that will be called if a message is published under this topic
	Handle func(i interface{}) error
	// Broker that this subscriber assigned to
	Broker *Broker
	// channel used to notify the end of subscription.
	exit chan bool
}

// Subscribe will attach this subscriber to its broker.
// Subscriber are not automaticaly attached when you get it from a broker.
// Until you call this function, this subscriber will not receive any message.
func (s *Subscriber) Subscribe() {
	s.Broker.lock.Lock()
	defer s.Broker.lock.Unlock()

	var prop observer.Property
	if p, ok := s.Broker.StreamMap[s.Topic]; ok {
		prop = p
	} else {
		prop = observer.NewProperty(nil)
		s.Broker.StreamMap[s.Topic] = prop
	}
	go s.observe(prop.Observe())
}

func (s *Subscriber) observe(stream observer.Stream) {
breakLoop:
	for {
		if err := s.Handle(stream.Value()); err != nil {
			fmt.Printf("Got error : %s\n", err)
		}
		select {
		case _, isOpen := <-s.exit:
			if isOpen {
				close(s.exit)
			}
			break breakLoop
		case <-stream.Changes():
			stream.Next()
		}
	}
}

// Unsubscribe will detach this subscriber from its broker.
func (s *Subscriber) Unsubscribe() {
	s.Broker.lock.Lock()
	defer s.Broker.lock.Unlock()

	if s.exit == nil {
		s.exit = make(chan bool)
	}
	s.exit <- true
	s.exit = nil
}

// Publisher is the publishing object.
// It automatically assigned with its broker and topic
type Publisher struct {
	Topic  string
	Broker *Broker
}

// Publish a message into its broker on its assigned topic.
func (p *Publisher) Publish(data interface{}) {
	p.Broker.lock.Lock()
	defer p.Broker.lock.Unlock()

	if prop, ok := p.Broker.StreamMap[p.Topic]; ok {
		prop.Update(data)
	}
}

// NewBroker creates new broker.
// If you don't need a specific broker, you can use the already created DefaultBroker
func NewBroker() *Broker {
	return &Broker{
		StreamMap: make(map[string]observer.Property),
	}
}

// Broker holds maps of topics.
type Broker struct {
	lock      sync.Mutex
	StreamMap map[string]observer.Property
}

// GetPublisher will obtain a Publisher dedicated for this broker and topic
func (b *Broker) GetPublisher(topic string) *Publisher {
	return &Publisher{
		Topic:  topic,
		Broker: b,
	}
}

// GetSubscriber will obtain  Subscriber dedicated for this broker and topic with a handler function that will
// handle the message that it's received.
// In the Handle function, implementation MUST check for nil argument.
func (b *Broker) GetSubscriber(topic string, Handle func(i interface{}) error) *Subscriber {
	return &Subscriber{
		Topic:  topic,
		Handle: Handle,
		Broker: b,
	}
}
