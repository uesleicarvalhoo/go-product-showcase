package crud

import "context"

func (c Crud[Model, Entity]) Create(ctx context.Context, e Entity) error {
	m := c.fromDomain(e)

	return c.db.WithContext(ctx).Create(&m).Error
}
