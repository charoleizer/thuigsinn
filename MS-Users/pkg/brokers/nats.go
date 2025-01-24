package brokers

import "github.com/nats-io/nats.go"

type Nats struct {
	JS nats.JetStreamContext
}

func NewNats(js nats.JetStreamContext) *Nats {
	return &Nats{
		JS: js,
	}
}

func (b *Nats) Publish(subject string, data []byte) error {
	_, err := b.JS.Publish(subject, data)
	if err != nil {
		return err
	}

	return nil
}
