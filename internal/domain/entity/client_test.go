package entity_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/dto"
)

func TestNewClientErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario  string
		payload   dto.CreateClientPayload
		errorMsgs []string
	}{
		{
			scenario: "when name is empty",
			payload: dto.CreateClientPayload{
				Email: gofakeit.Email(),
				Phone: gofakeit.Phone(),
				Address: dto.AddressPayload{
					ZipCode: gofakeit.Zip(),
					Street:  gofakeit.Street(),
					City:    gofakeit.City(),
				},
			},
			errorMsgs: []string{"Name is required"},
		},
		{
			scenario: "when name length is lower then 4",
			payload: dto.CreateClientPayload{
				Name:  "abc",
				Email: gofakeit.Email(),
				Phone: gofakeit.Phone(),
				Address: dto.AddressPayload{
					ZipCode: gofakeit.Zip(),
					Street:  gofakeit.Street(),
					City:    gofakeit.City(),
				},
			},
			errorMsgs: []string{"Name min length is 4"},
		},
		{
			scenario: "when email is empty",
			payload: dto.CreateClientPayload{
				Name:  gofakeit.Name(),
				Phone: gofakeit.Phone(),
				Address: dto.AddressPayload{
					ZipCode: gofakeit.Zip(),
					Street:  gofakeit.Street(),
					City:    gofakeit.City(),
				},
			},
			errorMsgs: []string{"Email is required"},
		},
		{
			scenario: "when email is invalid",
			payload: dto.CreateClientPayload{
				Name:  gofakeit.Name(),
				Email: "email@invalidcom",
				Phone: gofakeit.Phone(),
				Address: dto.AddressPayload{
					ZipCode: gofakeit.Zip(),
					Street:  gofakeit.Street(),
					City:    gofakeit.City(),
				},
			},
			errorMsgs: []string{"'Email' has value of 'email@invalidcom' which doesn't satisfy 'email'"},
		},
		{
			scenario: "when phone is empty",
			payload: dto.CreateClientPayload{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
				Address: dto.AddressPayload{
					ZipCode: gofakeit.Zip(),
					Street:  gofakeit.Street(),
					City:    gofakeit.City(),
				},
			},
			errorMsgs: []string{"Phone is required"},
		},
		{
			scenario: "when address is empty",
			payload: dto.CreateClientPayload{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
				Phone: gofakeit.Phone(),
			},
			errorMsgs: []string{
				"ZipCode is required",
				"Street is required",
				"City is required",
			},
		},
		{
			scenario: "when all fields are invalid should be return all errors",
			payload:  dto.CreateClientPayload{},
			errorMsgs: []string{
				"Name is required",
				"Email is required",
				"Phone is required",
				"ZipCode is required",
				"Street is required",
				"City is required",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()
			// Arrange
			payload := tc.payload

			// Action
			_, err := entity.NewClient(
				payload.Name, payload.Email, payload.Phone, payload.Address.ZipCode, payload.Address.Street, payload.Address.City,
			)

			// Assert
			assert.Error(t, err)
			for _, errMsg := range tc.errorMsgs {
				assert.ErrorContains(t, err, errMsg)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	t.Parallel()

	// Arrange
	payload := dto.CreateClientPayload{
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
		Phone: gofakeit.Phone(),
		Address: dto.AddressPayload{
			ZipCode: gofakeit.Zip(),
			Street:  gofakeit.Street(),
			City:    gofakeit.City(),
		},
	}

	// Action
	client, err := entity.NewClient(
		payload.Name, payload.Email, payload.Phone, payload.Address.ZipCode, payload.Address.Street, payload.Address.City,
	)

	// Assert
	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, client.ID)
	assert.Equal(t, payload.Name, client.Name)
	assert.Equal(t, payload.Email, client.Email)
	assert.Equal(t, payload.Phone, client.Phone)
	assert.Equal(t, payload.Address.City, client.Address.City)
	assert.Equal(t, payload.Address.Street, client.Address.Street)
	assert.Equal(t, payload.Address.ZipCode, client.Address.ZipCode)
}
