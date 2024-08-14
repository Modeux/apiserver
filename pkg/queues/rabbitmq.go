package queues

import (
	"context"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type RabbitMqInterface interface {
	GetConn() *amqp.Connection
	Publish(queueName string, payload []byte, priority uint8) error
	GetChannelV2(queueName string) (<-chan amqp.Delivery, *amqp.Channel, error)
}

type RabbitMq struct {
	Conn *amqp.Connection
}

func NewRabbitMq(c *amqp.Connection) RabbitMqInterface {
	return &RabbitMq{Conn: c}
}

func (r *RabbitMq) GetConn() *amqp.Connection {
	return r.Conn
}

func (r *RabbitMq) Publish(queueName string, payload []byte, priority uint8) error {
	ch, err := r.Conn.Channel()
	if err != nil {
		return errors.Wrap(err, "Publish")
	}
	defer ch.Close()

	args := make(amqp.Table)
	args["x-max-priority"] = uint8(9)
	q, err := ch.QueueDeclare(queueName, true, false, false, false, args)
	if err != nil {
		return errors.Wrap(err, "Publish")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        payload,
			Priority:    priority,
		})

	return nil
}

func (r *RabbitMq) GetChannelV2(queueName string) (<-chan amqp.Delivery, *amqp.Channel, error) {
	ch, err := r.Conn.Channel()
	if err != nil {
		return nil, nil, errors.Wrap(err, "GetChannel")
	}
	args := make(amqp.Table)
	args["x-max-priority"] = uint8(9)
	_, err = ch.QueueDeclare(queueName, true, false, false, false, args)
	if err != nil {
		return nil, nil, errors.Wrap(err, "PriorityQueue")
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		return nil, nil, errors.Wrap(err, "GetChannel")
	}
	msgs, err := ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		return nil, nil, err
	}
	return msgs, ch, nil
}
