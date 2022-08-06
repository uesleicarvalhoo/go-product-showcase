package database

import (
	"github.com/uesleicarvalhoo/go-product-showcase/pkg/database/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteMemory(Config) (*gorm.DB, error) {
	db, err := sqlite.NewMemoryConnection()
	if err != nil {
		return nil, err
	}

	if err := sqlite.AutoMigrate(db); err != nil {
		return nil, err
	}

	return db, nil
}
