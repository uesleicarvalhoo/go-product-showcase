package broker

import "github.com/uesleicarvalhoo/go-product-showcase/pkg/broker"

type MemoryBroker = broker.Memory

func NewMemory(Config) *broker.Memory {
	return broker.NewMemoryBroker()
}
