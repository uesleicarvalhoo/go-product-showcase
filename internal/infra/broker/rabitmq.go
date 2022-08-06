package broker

import "github.com/uesleicarvalhoo/go-product-showcase/pkg/broker"

func NewRabbitMQ(config Config) (*broker.RabbitMQClient, error) {
	return broker.NewRabbitMqClient(config.User, config.Password, config.Host, config.Port)
}
