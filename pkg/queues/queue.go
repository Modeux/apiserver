package queues

import (
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
)

type QueueInterface interface {
	GetRabbitMq() RabbitMqInterface
}

type Queue struct {
	Rabbitmq RabbitMqInterface
}

func NewQueue() (QueueInterface, error) {
	rConn, err := amqp.Dial(os.Getenv("RABBIT_MQ_HOST"))
	if err != nil {
		return nil, errors.Wrap(err, "NewQueue")
	}
	rq := NewRabbitMq(rConn)
	return &Queue{Rabbitmq: rq}, nil
}

func (q *Queue) GetRabbitMq() RabbitMqInterface {
	return q.Rabbitmq
}
