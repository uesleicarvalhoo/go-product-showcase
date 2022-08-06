package broker

import (
	"context"
	"sync"

	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
)

type Memory struct {
	sync.Mutex
	events map[string][]dto.Event
}

func NewMemoryBroker() *Memory {
	return &Memory{events: make(map[string][]dto.Event)}
}

func (m *Memory) SendEvent(ctx context.Context, event dto.Event) {
	m.Lock()
	defer m.Unlock()

	m.events[event.Topic] = append(m.events[event.Topic], event)
}

func (m *Memory) GetEvents(key string) []dto.Event {
	m.Lock()
	defer m.Unlock()

	return m.events[key]
}
