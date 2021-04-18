package mq

import (
	"github.com/bee-well/auth/config"
	"github.com/streadway/amqp"
)

type MqHandlerFunc func([]byte)

type Mq interface {
	Publish(string, []byte) error
	AttachHandler(string, MqHandlerFunc) error
}

var mock Mq

func NewMq() Mq {
	if mock != nil {
		return mock
	}
	return &rabbitMq{}
}

type rabbitMq struct {
	conn *amqp.Connection
}

func (r *rabbitMq) Publish(queue string, content []byte) error {
	conn, err := amqp.Dial(config.GetString(config.MqConnectionUrl))
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := r.declareQueue(ch, queue)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        content,
		},
	)
}

func (r *rabbitMq) AttachHandler(queue string, handler MqHandlerFunc) error {
	go r.attachHandler(queue, handler)
	return nil
}

func (r *rabbitMq) attachHandler(queue string, handler MqHandlerFunc) error {
	conn, err := amqp.Dial(config.GetString(config.MqConnectionUrl))
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := r.declareQueue(ch, queue)
	if err != nil {
		return err
	}

	msgs, _ := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			handler(d.Body)
		}
	}()
	<-forever

	return nil
}

func (rabbitMq) declareQueue(ch *amqp.Channel, queue string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)
}
