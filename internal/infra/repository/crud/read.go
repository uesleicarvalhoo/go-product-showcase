package crud

import (
	"context"

	"github.com/google/uuid"
)

func (c Crud[Model, Entity]) Fetch(ctx context.Context, id uuid.UUID) (Entity, error) {
	var m Model

	if tx := c.db.WithContext(ctx).First(&m, "id = ?", id); tx.Error != nil {
		var e Entity

		return e, tx.Error
	}

	return c.toDomain(m), nil
}

func (c Crud[Model, Entity]) FetchAll(ctx context.Context) ([]Entity, error) {
	var models []Model

	if tx := c.db.WithContext(ctx).Find(&models); tx.Error != nil {
		return nil, tx.Error
	}

	entities := make([]Entity, len(models))
	for i, e := range models {
		entities[i] = c.toDomain(e)
	}

	return entities, nil
}
