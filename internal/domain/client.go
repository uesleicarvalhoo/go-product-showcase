package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/usecase/client"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
)

type (
	Client           = entity.Client
	ClientAdress     = entity.ClientAddress
	ClientRepository = client.Repository
)

type ClientUseCase interface {
	Create(ctx context.Context, payload dto.CreateClientPayload) (Client, error)
	Fetch(ctx context.Context, id uuid.UUID) (Client, error)
	FetchAll(ctx context.Context) ([]Client, error)
	Update(ctx context.Context, id uuid.UUID, payload dto.UpdateClientPayload) (entity.Client, error)
}

func NewClientUseCase(r client.Repository, b client.Broker, eventTopic string) client.UseCase {
	return client.New(r, b, eventTopic)
}
