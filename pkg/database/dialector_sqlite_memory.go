package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDialectorSQLiteMemory() gorm.Dialector {
	return sqlite.Open("file::memory:")
	// return sqlite.Open("file:test_data.sqlite3")
}
