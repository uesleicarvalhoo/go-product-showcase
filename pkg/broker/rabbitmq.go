package broker

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/streadway/amqp"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/json"
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/logger"
)

const (
	maxConnectionRetries = 5
	retryConnectionDelay = 1
)

type RabbitMQClient struct {
	mu         sync.Mutex
	uri        string
	connection *amqp.Connection
	channel    *amqp.Channel
	errChannel chan *amqp.Error
}

func NewRabbitMqClient(user, password, host, port string) (*RabbitMQClient, error) {
	client := &RabbitMQClient{
		uri:        fmt.Sprintf("amqp://%s:%s@%s", user, password, net.JoinHostPort(host, port)),
		errChannel: make(chan *amqp.Error, 1),
	}

	if err := client.connect(); err != nil {
		return nil, err
	}

	return client, nil
}

func (mq *RabbitMQClient) Close() {
	mq.channel.Close()
	mq.connection.Close()
}

func (mq *RabbitMQClient) SendEvent(ctx context.Context, event dto.Event) {
	if err := mq.sendEvent(event); err != nil {
		logger.Errorf("failed to publish event: %s", err)
	}
}

func (mq *RabbitMQClient) sendEvent(event dto.Event) error {
	if err := mq.connect(); err != nil {
		return err
	}

	body, err := json.Encode(event.Data)
	if err != nil {
		return fmt.Errorf("failed to decode event data: %w", err)
	}

	err = mq.channel.Publish(event.Topic, event.Key, false, false, amqp.Publishing{
		Body: body,
	})
	if err != nil {
		if errors.Is(err, amqp.ErrClosed) {
			logger.Warningf("[RabbitMQ] Connection error, retrying to send event %+v", event)

			return mq.sendEvent(event)
		}

		return err
	}

	return nil
}

func (mq *RabbitMQClient) connect() error {
	if mq.connection != nil && !mq.connection.IsClosed() {
		return nil
	}

	con, err := amqp.Dial(mq.uri)
	if err != nil {
		return err
	}

	channel, err := con.Channel()
	if err != nil {
		return err
	}

	mq.errChannel = channel.NotifyClose(make(chan *amqp.Error))

	con.IsClosed()
	mq.connection = con
	mq.channel = channel

	go mq.handleConnectionError()

	return nil
}

func (mq *RabbitMQClient) handleConnectionError() {
	for range mq.errChannel {
		mq.mu.Lock()
		logger.Error("rabbitMQ connection is closed, trying stablish a new connection..")

		for i := 0; i < maxConnectionRetries; i++ {
			mq.connection = nil

			err := mq.connect()
			if err == nil {
				logger.Error("rabbitMQ connection re-established with success")

				break
			}

			logger.Errorf("failed to re-connect, '%s', trying again in %d seconds..", err, retryConnectionDelay)
			time.Sleep(time.Second * retryConnectionDelay)
		}

		if mq.channel == nil {
			logger.Panicf("Couldn't reconnect to RabbitMQ")
		}
		mq.mu.Unlock()
	}
}
