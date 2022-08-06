package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewMemoryConnection() (*gorm.DB, error) {
	dsn := "file:memdb1?mode=memory&cached-shared"

	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}
