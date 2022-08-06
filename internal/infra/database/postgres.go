package database

import (
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/database/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQL(cfg Config) (*gorm.DB, error) {
	db, err := postgres.NewConnection(cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port)
	if err != nil {
		return nil, err
	}

	if err := postgres.Migrate(db, cfg.Database); err != nil {
		return nil, err
	}

	return db, nil
}
