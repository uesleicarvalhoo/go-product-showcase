package crud

import "context"

func (c Crud[Model, Entity]) Update(ctx context.Context, e Entity) error {
	m := c.fromDomain(e)

	return c.db.WithContext(ctx).Save(m).Error
}
