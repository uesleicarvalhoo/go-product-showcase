package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uesleicarvalhoo/go-product-showcase/internal/domain/entity"
)

func TestValidateError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario       string
		model          interface{}
		expectedErrors []string
	}{
		{
			scenario: "when field has tag 'required' and field is empty",
			model: struct {
				Field string `validate:"required"`
			}{},
			expectedErrors: []string{"Field is required"},
		},
		{
			scenario: "when field has tag 'min' and field length is lower then tag value",
			model: struct {
				Field string `validate:"min=4"`
			}{Field: ""},
			expectedErrors: []string{"Field min length is 4"},
		},
		{
			scenario: "when field has tag 'max' and field length is higher then tag value",
			model: struct {
				Field string `validate:"max=10"`
			}{Field: "higher then 4"},
			expectedErrors: []string{"Field max length is 10"},
		},
		{
			scenario: "when field has an invalid tag should has default message",
			model: struct {
				Field string `validate:"email"`
			}{Field: "ueser@emailcom"},
			expectedErrors: []string{"'Field' has value of 'ueser@emailcom' which doesn't satisfy 'email'"},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()
			// Action
			err := entity.Validate(tc.model)

			// Assert
			for _, errMsg := range tc.expectedErrors {
				assert.ErrorContains(t, err, errMsg)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		model    interface{}
	}{
		{
			scenario: "when field has tag 'required' and field isn't empty",
			model: struct {
				Field string `validate:"required"`
			}{Field: "value"},
		},
		{
			scenario: "when field has tag 'min' and field length is higher then tag value",
			model: struct {
				Field string `validate:"min=4"`
			}{Field: "higher then 4"},
		},
		{
			scenario: "when field has tag 'max' and field length is lower then tag value",
			model: struct {
				Field string `validate:"max=10"`
			}{Field: "value"},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.scenario, func(t *testing.T) {
			t.Parallel()
			// Action
			err := entity.Validate(tc.model)

			// Assert
			assert.NoError(t, err)
		})
	}
}
