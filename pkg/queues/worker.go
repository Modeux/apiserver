package queues

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"video-heatmap/pkg/loggers"
)

type QueueWorkerInterface interface {
	Work(ctx context.Context, wg *sync.WaitGroup) error
}

type WorkerPool struct {
	Queues    QueueInterface
	QueueName string
	PoolSize  int
}

func NewWorkerPool(mq QueueInterface, queueName string, poolSize int) *WorkerPool {
	return &WorkerPool{mq, queueName, poolSize}
}

func (w *WorkerPool) Work(ctx context.Context, wg *sync.WaitGroup, task func(d amqp.Delivery)) error {
	msgCh, ch, err := w.Queues.GetRabbitMq().GetChannelV2(w.QueueName)
	if err != nil {
		loggers.Logger.Errorf("%+v", err)
		return err
	}
	for i := 0; i < w.PoolSize; i++ {
		wg.Add(1)
		go func(ctx context.Context, wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				select {
				case d := <-msgCh:
					task(d)
				case <-ctx.Done():
					if err := ch.Cancel("", true); err != nil {
						loggers.Logger.Errorf("AMQP channel cancel error: %s", err)
					}
					if err := ch.Close(); err != nil {
						loggers.Logger.Errorf("AMQP channel close error: %s", err)
					}
					return
				}
			}
		}(ctx, wg)
	}
	return nil
}
